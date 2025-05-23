/* tslint:disable */
/* eslint-disable */
/**
 * ContentSourcesBackend
 * The API for the repositories of the content sources that you can use to create and manage repositories between third-party applications and the [Red Hat Hybrid Cloud Console](https://console.redhat.com). With these repositories, you can build and deploy images using Image Builder for Cloud, on-Premise, and Edge. You can handle tasks, search for required RPMs, fetch a GPGKey from the URL, and list the features within applications. 
 *
 * The version of the OpenAPI document: v1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ApiTaskInfoResponse } from './ApiTaskInfoResponse';
import {
    ApiTaskInfoResponseFromJSON,
    ApiTaskInfoResponseFromJSONTyped,
    ApiTaskInfoResponseToJSON,
    ApiTaskInfoResponseToJSONTyped,
} from './ApiTaskInfoResponse';
import type { ApiSnapshotResponse } from './ApiSnapshotResponse';
import {
    ApiSnapshotResponseFromJSON,
    ApiSnapshotResponseFromJSONTyped,
    ApiSnapshotResponseToJSON,
    ApiSnapshotResponseToJSONTyped,
} from './ApiSnapshotResponse';

/**
 * 
 * @export
 * @interface ApiRepositoryResponse
 */
export interface ApiRepositoryResponse {
    /**
     * Account ID of the owner
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    readonly accountId?: string;
    /**
     * Content Type (rpm) of the repository
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    contentType?: string;
    /**
     * Architecture to restrict client usage to
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    distributionArch?: string;
    /**
     * Versions to restrict client usage to
     * @type {Array<string>}
     * @memberof ApiRepositoryResponse
     */
    distributionVersions?: Array<string>;
    /**
     * Number of consecutive failed introspections
     * @type {number}
     * @memberof ApiRepositoryResponse
     */
    failedIntrospectionsCount?: number;
    /**
     * Number of consecutive failed snapshots
     * @type {number}
     * @memberof ApiRepositoryResponse
     */
    failedSnapshotCount?: number;
    /**
     * The feature name this repo requires
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    featureName?: string;
    /**
     * GPG key for repository
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    gpgKey?: string;
    /**
     * Label used to configure the yum repository on clients
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    label?: string;
    /**
     * Error of last attempted introspection
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    lastIntrospectionError?: string;
    /**
     * Status of last introspection
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    lastIntrospectionStatus?: string;
    /**
     * Timestamp of last attempted introspection
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    lastIntrospectionTime?: string;
    /**
     * 
     * @type {ApiSnapshotResponse}
     * @memberof ApiRepositoryResponse
     */
    lastSnapshot?: ApiSnapshotResponse;
    /**
     * 
     * @type {ApiTaskInfoResponse}
     * @memberof ApiRepositoryResponse
     */
    lastSnapshotTask?: ApiTaskInfoResponse;
    /**
     * UUID of the last snapshot task
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    lastSnapshotTaskUuid?: string;
    /**
     * UUID of the last dao.Snapshot
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    lastSnapshotUuid?: string;
    /**
     * Timestamp of last successful introspection
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    lastSuccessIntrospectionTime?: string;
    /**
     * Timestamp of last introspection that had updates
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    lastUpdateIntrospectionTime?: string;
    /**
     * Latest URL for the snapshot distribution
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    latestSnapshotUrl?: string;
    /**
     * Verify packages
     * @type {boolean}
     * @memberof ApiRepositoryResponse
     */
    metadataVerification?: boolean;
    /**
     * Disable modularity filtering on this repository
     * @type {boolean}
     * @memberof ApiRepositoryResponse
     */
    moduleHotfixes?: boolean;
    /**
     * Name of the remote yum repository
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    name?: string;
    /**
     * Organization ID of the owner
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    readonly orgId?: string;
    /**
     * Origin of the repository
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    origin?: string;
    /**
     * Number of packages last read in the repository
     * @type {number}
     * @memberof ApiRepositoryResponse
     */
    packageCount?: number;
    /**
     * Enable snapshotting and hosting of this repository
     * @type {boolean}
     * @memberof ApiRepositoryResponse
     */
    snapshot?: boolean;
    /**
     * Combined status of last introspection and snapshot of repository (Valid, Invalid, Unavailable, Pending)
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    status?: string;
    /**
     * URL of the remote yum repository
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    url?: string;
    /**
     * UUID of the object
     * @type {string}
     * @memberof ApiRepositoryResponse
     */
    readonly uuid?: string;
}

/**
 * Check if a given object implements the ApiRepositoryResponse interface.
 */
export function instanceOfApiRepositoryResponse(value: object): value is ApiRepositoryResponse {
    return true;
}

export function ApiRepositoryResponseFromJSON(json: any): ApiRepositoryResponse {
    return ApiRepositoryResponseFromJSONTyped(json, false);
}

export function ApiRepositoryResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApiRepositoryResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'accountId': json['account_id'] == null ? undefined : json['account_id'],
        'contentType': json['content_type'] == null ? undefined : json['content_type'],
        'distributionArch': json['distribution_arch'] == null ? undefined : json['distribution_arch'],
        'distributionVersions': json['distribution_versions'] == null ? undefined : json['distribution_versions'],
        'failedIntrospectionsCount': json['failed_introspections_count'] == null ? undefined : json['failed_introspections_count'],
        'failedSnapshotCount': json['failed_snapshot_count'] == null ? undefined : json['failed_snapshot_count'],
        'featureName': json['feature_name'] == null ? undefined : json['feature_name'],
        'gpgKey': json['gpg_key'] == null ? undefined : json['gpg_key'],
        'label': json['label'] == null ? undefined : json['label'],
        'lastIntrospectionError': json['last_introspection_error'] == null ? undefined : json['last_introspection_error'],
        'lastIntrospectionStatus': json['last_introspection_status'] == null ? undefined : json['last_introspection_status'],
        'lastIntrospectionTime': json['last_introspection_time'] == null ? undefined : json['last_introspection_time'],
        'lastSnapshot': json['last_snapshot'] == null ? undefined : ApiSnapshotResponseFromJSON(json['last_snapshot']),
        'lastSnapshotTask': json['last_snapshot_task'] == null ? undefined : ApiTaskInfoResponseFromJSON(json['last_snapshot_task']),
        'lastSnapshotTaskUuid': json['last_snapshot_task_uuid'] == null ? undefined : json['last_snapshot_task_uuid'],
        'lastSnapshotUuid': json['last_snapshot_uuid'] == null ? undefined : json['last_snapshot_uuid'],
        'lastSuccessIntrospectionTime': json['last_success_introspection_time'] == null ? undefined : json['last_success_introspection_time'],
        'lastUpdateIntrospectionTime': json['last_update_introspection_time'] == null ? undefined : json['last_update_introspection_time'],
        'latestSnapshotUrl': json['latest_snapshot_url'] == null ? undefined : json['latest_snapshot_url'],
        'metadataVerification': json['metadata_verification'] == null ? undefined : json['metadata_verification'],
        'moduleHotfixes': json['module_hotfixes'] == null ? undefined : json['module_hotfixes'],
        'name': json['name'] == null ? undefined : json['name'],
        'orgId': json['org_id'] == null ? undefined : json['org_id'],
        'origin': json['origin'] == null ? undefined : json['origin'],
        'packageCount': json['package_count'] == null ? undefined : json['package_count'],
        'snapshot': json['snapshot'] == null ? undefined : json['snapshot'],
        'status': json['status'] == null ? undefined : json['status'],
        'url': json['url'] == null ? undefined : json['url'],
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
    };
}

export function ApiRepositoryResponseToJSON(json: any): ApiRepositoryResponse {
    return ApiRepositoryResponseToJSONTyped(json, false);
}

export function ApiRepositoryResponseToJSONTyped(value?: Omit<ApiRepositoryResponse, 'account_id'|'org_id'|'uuid'> | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'content_type': value['contentType'],
        'distribution_arch': value['distributionArch'],
        'distribution_versions': value['distributionVersions'],
        'failed_introspections_count': value['failedIntrospectionsCount'],
        'failed_snapshot_count': value['failedSnapshotCount'],
        'feature_name': value['featureName'],
        'gpg_key': value['gpgKey'],
        'label': value['label'],
        'last_introspection_error': value['lastIntrospectionError'],
        'last_introspection_status': value['lastIntrospectionStatus'],
        'last_introspection_time': value['lastIntrospectionTime'],
        'last_snapshot': ApiSnapshotResponseToJSON(value['lastSnapshot']),
        'last_snapshot_task': ApiTaskInfoResponseToJSON(value['lastSnapshotTask']),
        'last_snapshot_task_uuid': value['lastSnapshotTaskUuid'],
        'last_snapshot_uuid': value['lastSnapshotUuid'],
        'last_success_introspection_time': value['lastSuccessIntrospectionTime'],
        'last_update_introspection_time': value['lastUpdateIntrospectionTime'],
        'latest_snapshot_url': value['latestSnapshotUrl'],
        'metadata_verification': value['metadataVerification'],
        'module_hotfixes': value['moduleHotfixes'],
        'name': value['name'],
        'origin': value['origin'],
        'package_count': value['packageCount'],
        'snapshot': value['snapshot'],
        'status': value['status'],
        'url': value['url'],
    };
}

