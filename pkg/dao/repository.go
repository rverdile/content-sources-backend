package dao

import (
	"fmt"
	"strings"

	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type repositoryDaoImpl struct {
	db *gorm.DB
}

func GetRepositoryDao(db *gorm.DB) RepositoryDao {
	return repositoryDaoImpl{
		db: db,
	}
}

func DBErrorToApi(e error) *Error {
	pgError, ok := e.(*pgconn.PgError)
	if ok {
		if pgError.Code == "23505" {
			return &Error{BadValidation: true, Message: "Repository with this URL already belongs to organization"}
		}
	}
	dbError, ok := e.(models.Error)
	if ok {
		return &Error{BadValidation: dbError.Validation, Message: dbError.Message}
	}
	return &Error{Message: e.Error()}
}

func (r repositoryDaoImpl) Create(newRepoReq api.RepositoryRequest) (api.RepositoryResponse, error) {
	var newRepo models.Repository
	var newRepoConfig models.RepositoryConfiguration
	ApiFieldsToModel(newRepoReq, &newRepoConfig, &newRepo)

	if err := r.db.Where("url = ?", newRepo.URL).FirstOrCreate(&newRepo).Error; err != nil {
		return api.RepositoryResponse{}, DBErrorToApi(err)
	}

	if newRepoReq.OrgID != nil {
		newRepoConfig.OrgID = *newRepoReq.OrgID
	}
	if newRepoReq.AccountID != nil {
		newRepoConfig.AccountID = *newRepoReq.AccountID
	}
	newRepoConfig.RepositoryUUID = newRepo.Base.UUID

	if err := r.db.Create(&newRepoConfig).Error; err != nil {
		return api.RepositoryResponse{}, DBErrorToApi(err)
	}

	var created api.RepositoryResponse
	ModelToApiFields(newRepoConfig, &created)

	return created, nil
}

func (r repositoryDaoImpl) BulkCreate(newRepositories []api.RepositoryRequest) ([]api.RepositoryResponse, error) {
	var created = make([]api.RepositoryResponse, 0)
	var err error

	if err = r.db.Transaction(func(tx *gorm.DB) error {
		created, err = r.bulkCreate(tx, newRepositories)
		return err
	}); err != nil {
		return []api.RepositoryResponse{}, err
	}
	return created, err
}

func (r repositoryDaoImpl) bulkCreate(tx *gorm.DB, newRepositories []api.RepositoryRequest) ([]api.RepositoryResponse, error) {
	size := len(newRepositories)
	newRepoConfigs := make([]models.RepositoryConfiguration, size)
	newRepos := make([]models.Repository, size)

	for i := 0; i < size; i++ {
		newRepoConfigs[i] = models.RepositoryConfiguration{
			AccountID: *(newRepositories[i].AccountID),
			OrgID:     *(newRepositories[i].OrgID),
		}
		ApiFieldsToModel(newRepositories[i], &newRepoConfigs[i], &newRepos[i])
	}

	for i := 0; i < size; i++ {

		if err := tx.Where("url = ?", newRepos[i].URL).FirstOrCreate(&newRepos[i]).Error; err != nil {
			dbErr := DBErrorToApi(err)
			errContext := fmt.Sprintf("Error creating repository \"%s\"", newRepoConfigs[i].Name)
			dbErr.Wrap(errContext)
			return nil, dbErr
		}

		newRepoConfigs[i].RepositoryUUID = newRepos[i].UUID
		if err := tx.Create(&newRepoConfigs[i]).Error; err != nil {
			dbErr := DBErrorToApi(err)
			errContext := fmt.Sprintf("Error creating repository \"%s\"", newRepoConfigs[i].Name)
			dbErr.Wrap(errContext)
			return nil, dbErr
		}
	}

	created := convertToResponses(newRepoConfigs)
	return created, nil
}

func (r repositoryDaoImpl) List(
	OrgID string,
	pageData api.PaginationData,
	filterData api.FilterData,
) (api.RepositoryCollectionResponse, int64, error) {
	var totalRepos int64
	repoConfigs := make([]models.RepositoryConfiguration, 0)

	filteredDB := r.db

	filteredDB = filteredDB.Where("org_id = ?", OrgID)

	if filterData.AvailableForArch != "" {
		filteredDB = filteredDB.Where("arch = ? OR arch = ''", filterData.AvailableForArch)
	}
	if filterData.AvailableForVersion != "" {
		filteredDB = filteredDB.
			Where("? = any (versions) OR array_length(versions, 1) IS NULL", filterData.AvailableForVersion)
	}

	if filterData.Search != "" {
		containsSearch := "%" + filterData.Search + "%"
		filteredDB = filteredDB.Where("name LIKE ? OR url LIKE ?", containsSearch, containsSearch)
	}

	if filterData.Arch != "" {
		arches := strings.Split(filterData.Arch, ",")
		filteredDB = filteredDB.Where("arch IN ?", arches)
	}

	if filterData.Version != "" {
		versions := strings.Split(filterData.Version, ",")
		filteredDB = filteredDB.Where("? = any (versions)", versions[0])
		for i := 1; i < len(versions); i++ {
			filteredDB.Or("? = any (versions)", versions[i])
		}
	}

	filteredDB.Find(&repoConfigs).Count(&totalRepos)
	filteredDB.Preload("Repository").Limit(pageData.Limit).Offset(pageData.Offset).Find(&repoConfigs)

	if filteredDB.Error != nil {
		return api.RepositoryCollectionResponse{}, totalRepos, filteredDB.Error
	}
	repos := convertToResponses(repoConfigs)
	return api.RepositoryCollectionResponse{Data: repos}, totalRepos, nil
}

func (r repositoryDaoImpl) Fetch(orgID string, uuid string) (api.RepositoryResponse, error) {
	repo := api.RepositoryResponse{}
	repoConfig, err := r.fetchRepoConfig(orgID, uuid)
	if err != nil {
		return repo, err
	}
	ModelToApiFields(repoConfig, &repo)
	return repo, err
}

func (r repositoryDaoImpl) fetchRepoConfig(orgID string, uuid string) (models.RepositoryConfiguration, error) {
	found := models.RepositoryConfiguration{}
	result := r.db.
		Preload("Repository").
		Where("UUID = ? AND ORG_ID = ?", uuid, orgID).
		First(&found)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return found, &Error{NotFound: true, Message: "Could not find repository with UUID " + uuid}
		} else {
			return found, result.Error
		}
	}
	return found, nil
}

func (r repositoryDaoImpl) Update(orgID string, uuid string, repoParams api.RepositoryRequest) error {
	var repo models.Repository
	var repoConfig models.RepositoryConfiguration
	var err error

	repoConfig, err = r.fetchRepoConfig(orgID, uuid)
	if err != nil {
		return err
	}
	ApiFieldsToModel(repoParams, &repoConfig, &repo)

	// If URL is included in params, search for existing
	// Repository record, or create a new one.
	// Then replace existing Repository/RepoConfig association.
	if repoParams.URL != nil {
		err = r.db.FirstOrCreate(&repo, "url = ?", *repoParams.URL).Error
		if err != nil {
			return DBErrorToApi(err)
		}
		repoConfig.RepositoryUUID = repo.UUID
	}

	repoConfig.Repository = models.Repository{}
	if err := r.db.Save(&repoConfig).Error; err != nil {
		return DBErrorToApi(err)
	}
	return nil
}

func (r repositoryDaoImpl) Delete(orgID string, uuid string) error {
	repoConfig := models.RepositoryConfiguration{}
	if err := r.db.Where("UUID = ? AND ORG_ID = ?", uuid, orgID).First(&repoConfig).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &Error{NotFound: true, Message: "Could not find repository with UUID " + uuid}
		} else {
			return err
		}
	}
	if err := r.db.Delete(&repoConfig).Error; err != nil {
		return err
	}
	return nil
}

func ApiFieldsToModel(apiRepo api.RepositoryRequest, repoConfig *models.RepositoryConfiguration, repo *models.Repository) {
	if apiRepo.Name != nil {
		repoConfig.Name = *apiRepo.Name
	}
	if apiRepo.DistributionArch != nil {
		repoConfig.Arch = *apiRepo.DistributionArch
	}
	if apiRepo.DistributionVersions != nil {
		repoConfig.Versions = *apiRepo.DistributionVersions
	}
	if apiRepo.URL != nil {
		repo.URL = *apiRepo.URL
	}
}

func ModelToApiFields(repoConfig models.RepositoryConfiguration, apiRepo *api.RepositoryResponse) {
	apiRepo.UUID = repoConfig.UUID
	apiRepo.Name = repoConfig.Name
	apiRepo.DistributionVersions = repoConfig.Versions
	apiRepo.DistributionArch = repoConfig.Arch
	apiRepo.AccountID = repoConfig.AccountID
	apiRepo.OrgID = repoConfig.OrgID
	apiRepo.URL = repoConfig.Repository.URL
}

// Converts the database models to our response objects
func convertToResponses(repoConfigs []models.RepositoryConfiguration) []api.RepositoryResponse {
	repos := make([]api.RepositoryResponse, len(repoConfigs))
	for i := 0; i < len(repoConfigs); i++ {
		ModelToApiFields(repoConfigs[i], &repos[i])
	}
	return repos
}
