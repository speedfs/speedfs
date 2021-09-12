package storage

import (
	"context"

	"github.com/speedfs/speedfs/proto/storage"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) UploadFile(ctx context.Context, cmd *storage.UploadFileCommand) (*storage.UploadFileReply, error) {
	return nil, nil

}

func (s *Service) UploadAppenderFile(ctx context.Context, cmd *storage.UploadAppenderFileCommand) (*storage.UploadAppenderFileReply, error) {
	return nil, nil
}

func (s *Service) SetMetadata(ctx context.Context, cmd *storage.SetMetadataCommand) (*storage.SetMetadataReply, error) {
	return nil, nil
}

func (s *Service) GetMetadata(ctx context.Context, cmd *storage.GetMetadataCommand) (*storage.GetMetadataReply, error) {
	return nil, nil
}

func (s *Service) DownloadFile(ctx context.Context, cmd *storage.DownloadFileCommand) (*storage.DownloadFileReply, error) {
	return nil, nil
}

func (s *Service) AppendFile(ctx context.Context, cmd *storage.AppendFileCommand) (*storage.AppendFileReply, error) {
	return nil, nil
}

func (s *Service) ModifyFile(ctx context.Context, cmd *storage.ModifyFileCommand) (*storage.ModifyFileReply, error) {
	return nil, nil
}

func (s *Service) TruncateFile(ctx context.Context, cmd *storage.TruncateFileCommand) (*storage.TruncateFileReply, error) {
	return nil, nil
}

func (s *Service) RenameFile(ctx context.Context, cmd *storage.RenameFileCommand) (*storage.RenameFileReply, error) {
	return nil, nil
}
