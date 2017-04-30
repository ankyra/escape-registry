package gcs

import (
    "io"
    "github.com/ankyra/escape-registry/shared"
)

type GoogleCloudStorageBackend struct {}

func NewGoogleCloudStorageBackend() *GoogleCloudStorageBackend {
    return &GoogleCloudStorageBackend{}
}

func (ls *GoogleCloudStorageBackend) Upload(releaseId *shared.ReleaseId, pkg io.ReadSeeker) (string, error) {
//    def _store_file_in_gcs(self, gcs_credentials, gcs_storage_bucket, file, release):
//        if gcs_storage_bucket is None:
//            return "Undefined bucket for GCS storage backend", 500
//        if gcs_credentials is None:
//            return "Undefined credentials for GCS storage backend", 500
//
//        credentials = gcs_client.Credentials(gcs_credentials)
//        bucket = gcs_client.Bucket(name=gcs_storage_bucket, credentials=credentials)
//        archive = "/".join([release.releasetype.name, release.application.name, 
//                            release.get_release_id() + ".tgz"])
//        try:
//            with bucket.open(archive, "w") as f:
//                f.write(file.read())
//        except Exception, e:
//            return "Could not upload archive: %s" % str(e), 500
//        return "File uploaded " + archive, 200
    return "", nil
}

func (ls *GoogleCloudStorageBackend) Download(uri string) (io.ReadSeeker, error) {
    return nil, nil
}
