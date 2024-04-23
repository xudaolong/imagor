package awsconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xudaolong/imagor"
	"github.com/xudaolong/imagor/config"
	"github.com/xudaolong/imagor/storage/s3storage"
)

func TestS3Empty(t *testing.T) {
	srv := config.CreateServer([]string{}, WithAWS)
	app := srv.App.(*imagor.Imagor)
	assert.Equal(t, 1, len(app.Loaders))
	assert.Empty(t, app.Storages)
	assert.Empty(t, app.ResultStorages)
}

func TestS3Loader(t *testing.T) {
	srv := config.CreateServer([]string{
		"-aws-region", "asdf",
		"-aws-access-key-id", "asdf",
		"-aws-secret-access-key", "asdf",
		"-s3-endpoint", "asdfasdf",
		"-s3-force-path-style",
		"-s3-safe-chars", "!",

		"-s3-loader-bucket", "a",
		"-s3-loader-base-dir", "foo",
		"-s3-loader-path-prefix", "abcd",
	}, WithAWS)
	app := srv.App.(*imagor.Imagor)
	loader := app.Loaders[0].(*s3storage.S3Storage)
	assert.Equal(t, "a", loader.Bucket)
	assert.Equal(t, "/foo/", loader.BaseDir)
	assert.Equal(t, "/abcd/", loader.PathPrefix)
	assert.Equal(t, "!", loader.SafeChars)
}

func TestS3Storage(t *testing.T) {
	srv := config.CreateServer([]string{
		"-aws-region", "asdf",
		"-aws-access-key-id", "asdf",
		"-aws-secret-access-key", "asdf",
		"-s3-endpoint", "asdfasdf",
		"-s3-force-path-style",
		"-s3-safe-chars", "!",

		"-s3-storage-bucket", "a",
		"-s3-storage-base-dir", "foo",
		"-s3-storage-path-prefix", "abcd",

		"-s3-result-storage-bucket", "b",
		"-s3-result-storage-base-dir", "bar",
		"-s3-result-storage-path-prefix", "bcda",
	}, WithAWS)
	app := srv.App.(*imagor.Imagor)
	assert.Equal(t, 1, len(app.Loaders))
	storage := app.Storages[0].(*s3storage.S3Storage)
	assert.Equal(t, "a", storage.Bucket)
	assert.Equal(t, "/foo/", storage.BaseDir)
	assert.Equal(t, "/abcd/", storage.PathPrefix)
	assert.Equal(t, "!", storage.SafeChars)

	resultStorage := app.ResultStorages[0].(*s3storage.S3Storage)
	assert.Equal(t, "b", resultStorage.Bucket)
	assert.Equal(t, "/bar/", resultStorage.BaseDir)
	assert.Equal(t, "/bcda/", resultStorage.PathPrefix)
	assert.Equal(t, "!", resultStorage.SafeChars)
}

func TestS3SessionOverride(t *testing.T) {
	srv := config.CreateServer([]string{
		"-aws-loader-region", "asdf",
		"-aws-loader-access-key-id", "asdf",
		"-aws-loader-secret-access-key", "asdf",
		"-s3-loader-endpoint", "asdfasdf",

		"-aws-storage-region", "ghkj",
		"-aws-storage-access-key-id", "ghkj",
		"-aws-storage-secret-access-key", "ghkj",
		"-s3-storage-endpoint", "asdfasdf",

		"-aws-result-storage-region", "qwer",
		"-aws-result-storage-access-key-id", "qwer",
		"-aws-result-storage-secret-access-key", "qwer",
		"-s3-result-storage-endpoint", "asdfasdf",

		"-s3-force-path-style",
		"-s3-safe-chars", "!",

		"-s3-loader-bucket", "a",
		"-s3-loader-base-dir", "foo",
		"-s3-loader-path-prefix", "abcd",

		"-s3-storage-bucket", "a",
		"-s3-storage-base-dir", "foo",
		"-s3-storage-path-prefix", "abcd",

		"-s3-result-storage-bucket", "b",
		"-s3-result-storage-base-dir", "bar",
		"-s3-result-storage-path-prefix", "bcda",
	}, WithAWS)
	app := srv.App.(*imagor.Imagor)
	loader := app.Loaders[0].(*s3storage.S3Storage)
	assert.Equal(t, "a", loader.Bucket)
	assert.Equal(t, "/foo/", loader.BaseDir)
	assert.Equal(t, "/abcd/", loader.PathPrefix)
	assert.Equal(t, "!", loader.SafeChars)

	storage := app.Storages[0].(*s3storage.S3Storage)
	assert.Equal(t, "a", storage.Bucket)
	assert.Equal(t, "/foo/", storage.BaseDir)
	assert.Equal(t, "/abcd/", storage.PathPrefix)
	assert.Equal(t, "!", storage.SafeChars)

	resultStorage := app.ResultStorages[0].(*s3storage.S3Storage)
	assert.Equal(t, "b", resultStorage.Bucket)
	assert.Equal(t, "/bar/", resultStorage.BaseDir)
	assert.Equal(t, "/bcda/", resultStorage.PathPrefix)
	assert.Equal(t, "!", resultStorage.SafeChars)
}

func TestS3StorageClassWithResultStorageBucket(t *testing.T) {
	srv := config.CreateServer([]string{
		"-s3-storage-class", "asdf",
		"-s3-storage-bucket", "a",
		"-s3-result-storage-bucket", "b",
	}, WithAWS)
	app := srv.App.(*imagor.Imagor)
	storage := app.Storages[0].(*s3storage.S3Storage)
	assert.Equal(t, "STANDARD", storage.StorageClass)

}

func TestS3StorageClassWithoutResultStorageBucket(t *testing.T) {
	srv := config.CreateServer([]string{
		"-s3-storage-class", "asdf",
		"-s3-storage-bucket", "a",
	}, WithAWS)
	app := srv.App.(*imagor.Imagor)
	storage := app.Storages[0].(*s3storage.S3Storage)
	assert.Equal(t, "STANDARD", storage.StorageClass)

}

func TestS3StorageClass(t *testing.T) {
	srv := config.CreateServer([]string{
		"-s3-storage-class", "asdf",
		"-s3-storage-bucket", "a",
	}, WithAWS)
	app := srv.App.(*imagor.Imagor)
	storage := app.Storages[0].(*s3storage.S3Storage)
	assert.Equal(t, "STANDARD", storage.StorageClass)

	srv = config.CreateServer([]string{
		"-s3-storage-class", "REDUCED_REDUNDANCY",
		"-s3-storage-bucket", "a",
	}, WithAWS)
	app = srv.App.(*imagor.Imagor)
	storage = app.Storages[0].(*s3storage.S3Storage)
	assert.Equal(t, "REDUCED_REDUNDANCY", storage.StorageClass)
}
