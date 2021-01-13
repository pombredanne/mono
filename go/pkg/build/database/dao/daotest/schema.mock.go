// Generated by protoc-ddl.
// protoc-gen-dao-mock: v0.1
package daotest

import (
	"context"
	"database/sql"

	"go.f110.dev/protoc-ddl/mock"

	"go.f110.dev/mono/go/pkg/build/database"
	"go.f110.dev/mono/go/pkg/build/database/dao"
)

type SourceRepository struct {
	*mock.Mock
}

func NewSourceRepository() *SourceRepository {
	return &SourceRepository{Mock: mock.New()}
}

func (d *SourceRepository) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	return nil

}

func (d *SourceRepository) Select(ctx context.Context, id int32) (*database.SourceRepository, error) {
	v, err := d.Call("Select", map[string]interface{}{"id": id})
	return v.(*database.SourceRepository), err

}

func (d *SourceRepository) RegisterSelect(id int32, value *database.SourceRepository) {
	d.Register("Select", map[string]interface{}{"id": id}, value, nil)
}

func (d *SourceRepository) ListAll(ctx context.Context, opt ...dao.ListOption) ([]*database.SourceRepository, error) {
	v, err := d.Call("ListAll", map[string]interface{}{})
	return v.([]*database.SourceRepository), err

}

func (d *SourceRepository) RegisterListAll(value []*database.SourceRepository, err error) {
	d.Register("ListAll", map[string]interface{}{}, value, err)
}

func (d *SourceRepository) ListByUrl(ctx context.Context, url string, opt ...dao.ListOption) ([]*database.SourceRepository, error) {
	v, err := d.Call("ListByUrl", map[string]interface{}{"url": url})
	return v.([]*database.SourceRepository), err

}

func (d *SourceRepository) RegisterListByUrl(url string, value []*database.SourceRepository, err error) {
	d.Register("ListByUrl", map[string]interface{}{"url": url}, value, err)
}

func (d *SourceRepository) Create(ctx context.Context, sourceRepository *database.SourceRepository, opt ...dao.ExecOption) (*database.SourceRepository, error) {
	_, _ = d.Call("Create", map[string]interface{}{"sourceRepository": sourceRepository})
	return sourceRepository, nil

}

func (d *SourceRepository) Delete(ctx context.Context, id int32, opt ...dao.ExecOption) error {
	_, _ = d.Call("Delete", map[string]interface{}{"id": id})
	return nil

}

func (d *SourceRepository) Update(ctx context.Context, sourceRepository *database.SourceRepository, opt ...dao.ExecOption) error {
	_, _ = d.Call("Update", map[string]interface{}{"sourceRepository": sourceRepository})
	return nil

}

type Job struct {
	*mock.Mock
}

func NewJob() *Job {
	return &Job{Mock: mock.New()}
}

func (d *Job) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	return nil

}

func (d *Job) Select(ctx context.Context, id int32) (*database.Job, error) {
	v, err := d.Call("Select", map[string]interface{}{"id": id})
	return v.(*database.Job), err

}

func (d *Job) RegisterSelect(id int32, value *database.Job) {
	d.Register("Select", map[string]interface{}{"id": id}, value, nil)
}

func (d *Job) ListAll(ctx context.Context, opt ...dao.ListOption) ([]*database.Job, error) {
	v, err := d.Call("ListAll", map[string]interface{}{})
	return v.([]*database.Job), err

}

func (d *Job) RegisterListAll(value []*database.Job, err error) {
	d.Register("ListAll", map[string]interface{}{}, value, err)
}

func (d *Job) ListBySourceRepositoryId(ctx context.Context, repositoryId int32, opt ...dao.ListOption) ([]*database.Job, error) {
	v, err := d.Call("ListBySourceRepositoryId", map[string]interface{}{"repositoryId": repositoryId})
	return v.([]*database.Job), err

}

func (d *Job) RegisterListBySourceRepositoryId(repositoryId int32, value []*database.Job, err error) {
	d.Register("ListBySourceRepositoryId", map[string]interface{}{"repositoryId": repositoryId}, value, err)
}

func (d *Job) Create(ctx context.Context, job *database.Job, opt ...dao.ExecOption) (*database.Job, error) {
	_, _ = d.Call("Create", map[string]interface{}{"job": job})
	return job, nil

}

func (d *Job) Delete(ctx context.Context, id int32, opt ...dao.ExecOption) error {
	_, _ = d.Call("Delete", map[string]interface{}{"id": id})
	return nil

}

func (d *Job) Update(ctx context.Context, job *database.Job, opt ...dao.ExecOption) error {
	_, _ = d.Call("Update", map[string]interface{}{"job": job})
	return nil

}

type Task struct {
	*mock.Mock
}

func NewTask() *Task {
	return &Task{Mock: mock.New()}
}

func (d *Task) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	return nil

}

func (d *Task) Select(ctx context.Context, id int32) (*database.Task, error) {
	v, err := d.Call("Select", map[string]interface{}{"id": id})
	return v.(*database.Task), err

}

func (d *Task) RegisterSelect(id int32, value *database.Task) {
	d.Register("Select", map[string]interface{}{"id": id}, value, nil)
}

func (d *Task) ListByJobId(ctx context.Context, jobId int32, opt ...dao.ListOption) ([]*database.Task, error) {
	v, err := d.Call("ListByJobId", map[string]interface{}{"jobId": jobId})
	return v.([]*database.Task), err

}

func (d *Task) RegisterListByJobId(jobId int32, value []*database.Task, err error) {
	d.Register("ListByJobId", map[string]interface{}{"jobId": jobId}, value, err)
}

func (d *Task) ListPending(ctx context.Context, opt ...dao.ListOption) ([]*database.Task, error) {
	v, err := d.Call("ListPending", map[string]interface{}{})
	return v.([]*database.Task), err

}

func (d *Task) RegisterListPending(value []*database.Task, err error) {
	d.Register("ListPending", map[string]interface{}{}, value, err)
}

func (d *Task) Create(ctx context.Context, task *database.Task, opt ...dao.ExecOption) (*database.Task, error) {
	_, _ = d.Call("Create", map[string]interface{}{"task": task})
	return task, nil

}

func (d *Task) Delete(ctx context.Context, id int32, opt ...dao.ExecOption) error {
	_, _ = d.Call("Delete", map[string]interface{}{"id": id})
	return nil

}

func (d *Task) Update(ctx context.Context, task *database.Task, opt ...dao.ExecOption) error {
	_, _ = d.Call("Update", map[string]interface{}{"task": task})
	return nil

}

type TrustedUser struct {
	*mock.Mock
}

func NewTrustedUser() *TrustedUser {
	return &TrustedUser{Mock: mock.New()}
}

func (d *TrustedUser) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	return nil

}

func (d *TrustedUser) Select(ctx context.Context, id int32) (*database.TrustedUser, error) {
	v, err := d.Call("Select", map[string]interface{}{"id": id})
	return v.(*database.TrustedUser), err

}

func (d *TrustedUser) RegisterSelect(id int32, value *database.TrustedUser) {
	d.Register("Select", map[string]interface{}{"id": id}, value, nil)
}

func (d *TrustedUser) ListAll(ctx context.Context, opt ...dao.ListOption) ([]*database.TrustedUser, error) {
	v, err := d.Call("ListAll", map[string]interface{}{})
	return v.([]*database.TrustedUser), err

}

func (d *TrustedUser) RegisterListAll(value []*database.TrustedUser, err error) {
	d.Register("ListAll", map[string]interface{}{}, value, err)
}

func (d *TrustedUser) ListByGithubId(ctx context.Context, githubId int64, opt ...dao.ListOption) ([]*database.TrustedUser, error) {
	v, err := d.Call("ListByGithubId", map[string]interface{}{"githubId": githubId})
	return v.([]*database.TrustedUser), err

}

func (d *TrustedUser) RegisterListByGithubId(githubId int64, value []*database.TrustedUser, err error) {
	d.Register("ListByGithubId", map[string]interface{}{"githubId": githubId}, value, err)
}

func (d *TrustedUser) Create(ctx context.Context, trustedUser *database.TrustedUser, opt ...dao.ExecOption) (*database.TrustedUser, error) {
	_, _ = d.Call("Create", map[string]interface{}{"trustedUser": trustedUser})
	return trustedUser, nil

}

func (d *TrustedUser) Delete(ctx context.Context, id int32, opt ...dao.ExecOption) error {
	_, _ = d.Call("Delete", map[string]interface{}{"id": id})
	return nil

}

func (d *TrustedUser) Update(ctx context.Context, trustedUser *database.TrustedUser, opt ...dao.ExecOption) error {
	_, _ = d.Call("Update", map[string]interface{}{"trustedUser": trustedUser})
	return nil

}

type PermitPullRequest struct {
	*mock.Mock
}

func NewPermitPullRequest() *PermitPullRequest {
	return &PermitPullRequest{Mock: mock.New()}
}

func (d *PermitPullRequest) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	return nil

}

func (d *PermitPullRequest) Select(ctx context.Context, id int32) (*database.PermitPullRequest, error) {
	v, err := d.Call("Select", map[string]interface{}{"id": id})
	return v.(*database.PermitPullRequest), err

}

func (d *PermitPullRequest) RegisterSelect(id int32, value *database.PermitPullRequest) {
	d.Register("Select", map[string]interface{}{"id": id}, value, nil)
}

func (d *PermitPullRequest) ListByRepositoryAndNumber(ctx context.Context, repository string, number int32, opt ...dao.ListOption) ([]*database.PermitPullRequest, error) {
	v, err := d.Call("ListByRepositoryAndNumber", map[string]interface{}{"repository": repository, "number": number})
	return v.([]*database.PermitPullRequest), err

}

func (d *PermitPullRequest) RegisterListByRepositoryAndNumber(repository string, number int32, value []*database.PermitPullRequest, err error) {
	d.Register("ListByRepositoryAndNumber", map[string]interface{}{"repository": repository, "number": number}, value, err)
}

func (d *PermitPullRequest) Create(ctx context.Context, permitPullRequest *database.PermitPullRequest, opt ...dao.ExecOption) (*database.PermitPullRequest, error) {
	_, _ = d.Call("Create", map[string]interface{}{"permitPullRequest": permitPullRequest})
	return permitPullRequest, nil

}

func (d *PermitPullRequest) Delete(ctx context.Context, id int32, opt ...dao.ExecOption) error {
	_, _ = d.Call("Delete", map[string]interface{}{"id": id})
	return nil

}

func (d *PermitPullRequest) Update(ctx context.Context, permitPullRequest *database.PermitPullRequest, opt ...dao.ExecOption) error {
	_, _ = d.Call("Update", map[string]interface{}{"permitPullRequest": permitPullRequest})
	return nil

}
