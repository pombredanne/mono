// Generated by protoc-ddl.
// protoc-gen-dao: v0.1
package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"golang.org/x/xerrors"

	"go.f110.dev/mono/tools/build/pkg/database"
)

type ListOption func(opt *listOpt)

func Limit(limit int) func(opt *listOpt) {
	return func(opt *listOpt) {
		opt.limit = limit
	}
}

func Desc(opt *listOpt) {
	opt.desc = true
}

type listOpt struct {
	limit int
	desc  bool
}

func newListOpt(opts ...ListOption) *listOpt {
	opt := &listOpt{}
	for _, v := range opts {
		v(opt)
	}
	return opt
}

type ExecOption func(opt *execOpt)

func WithTx(tx *sql.Tx) ExecOption {
	return func(opt *execOpt) {
		opt.tx = tx
	}
}

type execOpt struct {
	tx *sql.Tx
}

func newExecOpt(opts ...ExecOption) *execOpt {
	opt := &execOpt{}
	for _, v := range opts {
		v(opt)
	}
	return opt
}

type execConn interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type SourceRepository struct {
	conn *sql.DB
}

type SourceRepositoryInterface interface {
	Tx(ctx context.Context, fn func(tx *sql.Tx) error) error
	Select(ctx context.Context, id int32) (*database.SourceRepository, error)
	ListAll(ctx context.Context, opt ...ListOption) ([]*database.SourceRepository, error)
	ListByUrl(ctx context.Context, url string, opt ...ListOption) ([]*database.SourceRepository, error)
	Create(ctx context.Context, sourceRepository *database.SourceRepository, opt ...ExecOption) (*database.SourceRepository, error)
	Update(ctx context.Context, sourceRepository *database.SourceRepository, opt ...ExecOption) error
	Delete(ctx context.Context, id int32, opt ...ExecOption) error
}

var _ SourceRepositoryInterface = &SourceRepository{}

func NewSourceRepository(conn *sql.DB) *SourceRepository {
	return &SourceRepository{
		conn: conn,
	}
}

func (d *SourceRepository) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if err := fn(tx); err != nil {
		rErr := tx.Rollback()
		return xerrors.Errorf("%v: %w", rErr, err)
	}

	err = tx.Commit()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil

}

func (d *SourceRepository) Select(ctx context.Context, id int32) (*database.SourceRepository, error) {
	row := d.conn.QueryRowContext(ctx, "SELECT * FROM `source_repository` WHERE `id` = ?", id)

	v := &database.SourceRepository{}
	if err := row.Scan(&v.Id, &v.Url, &v.CloneUrl, &v.Name, &v.Private, &v.CreatedAt, &v.UpdatedAt); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	v.ResetMark()
	return v, nil

}

func (d *SourceRepository) ListAll(ctx context.Context, opt ...ListOption) ([]*database.SourceRepository, error) {
	listOpts := newListOpt(opt...)
	query := "select id, url, clone_url, name, private, created_at, updated_at from source_repository"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.SourceRepository, 0)
	for rows.Next() {
		r := &database.SourceRepository{}
		if err := rows.Scan(&r.Id, &r.Url, &r.CloneUrl, &r.Name, &r.Private, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}

	return res, nil

}

func (d *SourceRepository) ListByUrl(ctx context.Context, url string, opt ...ListOption) ([]*database.SourceRepository, error) {
	listOpts := newListOpt(opt...)
	query := "select id, url, clone_url, name, private, created_at, updated_at from source_repository where url = ?"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
		url,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.SourceRepository, 0)
	for rows.Next() {
		r := &database.SourceRepository{}
		if err := rows.Scan(&r.Id, &r.Url, &r.CloneUrl, &r.Name, &r.Private, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}

	return res, nil

}

func (d *SourceRepository) Create(ctx context.Context, sourceRepository *database.SourceRepository, opt ...ExecOption) (*database.SourceRepository, error) {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(
		ctx,
		"INSERT INTO `source_repository` (`url`, `clone_url`, `name`, `private`, `created_at`) VALUES (?, ?, ?, ?, ?)",
		sourceRepository.Url, sourceRepository.CloneUrl, sourceRepository.Name, sourceRepository.Private, time.Now(),
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return nil, sql.ErrNoRows
	}

	sourceRepository = sourceRepository.Copy()
	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	sourceRepository.Id = int32(insertedId)

	sourceRepository.ResetMark()
	return sourceRepository, nil

}

func (d *SourceRepository) Delete(ctx context.Context, id int32, opt ...ExecOption) error {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(ctx, "DELETE FROM `source_repository` WHERE `id` = ?", id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (d *SourceRepository) Update(ctx context.Context, sourceRepository *database.SourceRepository, opt ...ExecOption) error {
	if !sourceRepository.IsChanged() {
		return nil
	}

	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	changedColumn := sourceRepository.ChangedColumn()
	cols := make([]string, len(changedColumn)+1)
	values := make([]interface{}, len(changedColumn)+1)
	for i := range changedColumn {
		cols[i] = "`" + changedColumn[i].Name + "` = ?"
		values[i] = changedColumn[i].Value
	}
	cols[len(cols)-1] = "`updated_at` = ?"
	values[len(values)-1] = time.Now()

	query := fmt.Sprintf("UPDATE `source_repository` SET %s WHERE `id` = ?", strings.Join(cols, ", "))
	res, err := conn.ExecContext(
		ctx,
		query,
		append(values, sourceRepository.Id)...,
	)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	sourceRepository.ResetMark()
	return nil

}

type Job struct {
	conn *sql.DB

	sourceRepository *SourceRepository
}

type JobInterface interface {
	Tx(ctx context.Context, fn func(tx *sql.Tx) error) error
	Select(ctx context.Context, id int32) (*database.Job, error)
	ListAll(ctx context.Context, opt ...ListOption) ([]*database.Job, error)
	ListBySourceRepositoryId(ctx context.Context, repositoryId int32, opt ...ListOption) ([]*database.Job, error)
	Create(ctx context.Context, job *database.Job, opt ...ExecOption) (*database.Job, error)
	Update(ctx context.Context, job *database.Job, opt ...ExecOption) error
	Delete(ctx context.Context, id int32, opt ...ExecOption) error
}

var _ JobInterface = &Job{}

func NewJob(conn *sql.DB) *Job {
	return &Job{
		conn:             conn,
		sourceRepository: NewSourceRepository(conn),
	}
}

func (d *Job) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if err := fn(tx); err != nil {
		rErr := tx.Rollback()
		return xerrors.Errorf("%v: %w", rErr, err)
	}

	err = tx.Commit()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil

}

func (d *Job) Select(ctx context.Context, id int32) (*database.Job, error) {
	row := d.conn.QueryRowContext(ctx, "SELECT * FROM `job` WHERE `id` = ?", id)

	v := &database.Job{}
	if err := row.Scan(&v.Id, &v.RepositoryId, &v.Command, &v.Target, &v.Active, &v.AllRevision, &v.GithubStatus, &v.CpuLimit, &v.MemoryLimit, &v.Exclusive, &v.Sync, &v.ConfigName, &v.BazelVersion, &v.JobType, &v.CreatedAt, &v.UpdatedAt); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	{
		rel, err := d.sourceRepository.Select(ctx, v.RepositoryId)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		v.Repository = rel
	}

	v.ResetMark()
	return v, nil

}

func (d *Job) ListAll(ctx context.Context, opt ...ListOption) ([]*database.Job, error) {
	listOpts := newListOpt(opt...)
	query := "select id, repository_id, command, target, active, all_revision, github_status, cpu_limit, memory_limit, exclusive, sync, config_name, bazel_version, job_type, created_at, updated_at from job"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.Job, 0)
	for rows.Next() {
		r := &database.Job{}
		if err := rows.Scan(&r.Id, &r.RepositoryId, &r.Command, &r.Target, &r.Active, &r.AllRevision, &r.GithubStatus, &r.CpuLimit, &r.MemoryLimit, &r.Exclusive, &r.Sync, &r.ConfigName, &r.BazelVersion, &r.JobType, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}
	if len(res) > 0 {
		for _, v := range res {
			{
				rel, err := d.sourceRepository.Select(ctx, v.RepositoryId)
				if err != nil {
					return nil, xerrors.Errorf(": %w", err)
				}
				v.Repository = rel
			}

		}
	}

	return res, nil

}

func (d *Job) ListBySourceRepositoryId(ctx context.Context, repositoryId int32, opt ...ListOption) ([]*database.Job, error) {
	listOpts := newListOpt(opt...)
	query := "select id, repository_id, command, target, active, all_revision, github_status, cpu_limit, memory_limit, exclusive, sync, config_name, bazel_version, job_type, created_at, updated_at from job where repository_id = ?"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
		repositoryId,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.Job, 0)
	for rows.Next() {
		r := &database.Job{}
		if err := rows.Scan(&r.Id, &r.RepositoryId, &r.Command, &r.Target, &r.Active, &r.AllRevision, &r.GithubStatus, &r.CpuLimit, &r.MemoryLimit, &r.Exclusive, &r.Sync, &r.ConfigName, &r.BazelVersion, &r.JobType, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}
	if len(res) > 0 {
		for _, v := range res {
			{
				rel, err := d.sourceRepository.Select(ctx, v.RepositoryId)
				if err != nil {
					return nil, xerrors.Errorf(": %w", err)
				}
				v.Repository = rel
			}

		}
	}

	return res, nil

}

func (d *Job) Create(ctx context.Context, job *database.Job, opt ...ExecOption) (*database.Job, error) {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(
		ctx,
		"INSERT INTO `job` (`repository_id`, `command`, `target`, `active`, `all_revision`, `github_status`, `cpu_limit`, `memory_limit`, `exclusive`, `sync`, `config_name`, `bazel_version`, `job_type`, `created_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		job.RepositoryId, job.Command, job.Target, job.Active, job.AllRevision, job.GithubStatus, job.CpuLimit, job.MemoryLimit, job.Exclusive, job.Sync, job.ConfigName, job.BazelVersion, job.JobType, time.Now(),
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return nil, sql.ErrNoRows
	}

	job = job.Copy()
	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	job.Id = int32(insertedId)

	job.ResetMark()
	return job, nil

}

func (d *Job) Delete(ctx context.Context, id int32, opt ...ExecOption) error {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(ctx, "DELETE FROM `job` WHERE `id` = ?", id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (d *Job) Update(ctx context.Context, job *database.Job, opt ...ExecOption) error {
	if !job.IsChanged() {
		return nil
	}

	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	changedColumn := job.ChangedColumn()
	cols := make([]string, len(changedColumn)+1)
	values := make([]interface{}, len(changedColumn)+1)
	for i := range changedColumn {
		cols[i] = "`" + changedColumn[i].Name + "` = ?"
		values[i] = changedColumn[i].Value
	}
	cols[len(cols)-1] = "`updated_at` = ?"
	values[len(values)-1] = time.Now()

	query := fmt.Sprintf("UPDATE `job` SET %s WHERE `id` = ?", strings.Join(cols, ", "))
	res, err := conn.ExecContext(
		ctx,
		query,
		append(values, job.Id)...,
	)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	job.ResetMark()
	return nil

}

type Task struct {
	conn *sql.DB

	job *Job
}

type TaskInterface interface {
	Tx(ctx context.Context, fn func(tx *sql.Tx) error) error
	Select(ctx context.Context, id int32) (*database.Task, error)
	ListByJobId(ctx context.Context, jobId int32, opt ...ListOption) ([]*database.Task, error)
	ListPending(ctx context.Context, opt ...ListOption) ([]*database.Task, error)
	Create(ctx context.Context, task *database.Task, opt ...ExecOption) (*database.Task, error)
	Update(ctx context.Context, task *database.Task, opt ...ExecOption) error
	Delete(ctx context.Context, id int32, opt ...ExecOption) error
}

var _ TaskInterface = &Task{}

func NewTask(conn *sql.DB) *Task {
	return &Task{
		conn: conn,
		job:  NewJob(conn),
	}
}

func (d *Task) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if err := fn(tx); err != nil {
		rErr := tx.Rollback()
		return xerrors.Errorf("%v: %w", rErr, err)
	}

	err = tx.Commit()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil

}

func (d *Task) Select(ctx context.Context, id int32) (*database.Task, error) {
	row := d.conn.QueryRowContext(ctx, "SELECT * FROM `task` WHERE `id` = ?", id)

	v := &database.Task{}
	if err := row.Scan(&v.Id, &v.JobId, &v.Revision, &v.Success, &v.LogFile, &v.Command, &v.Target, &v.Via, &v.ConfigName, &v.StartAt, &v.FinishedAt, &v.CreatedAt, &v.UpdatedAt); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	{
		rel, err := d.job.Select(ctx, v.JobId)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		v.Job = rel
	}

	v.ResetMark()
	return v, nil

}

func (d *Task) ListByJobId(ctx context.Context, jobId int32, opt ...ListOption) ([]*database.Task, error) {
	listOpts := newListOpt(opt...)
	query := "select id, job_id, revision, success, log_file, command, target, via, config_name, start_at, finished_at, created_at, updated_at from task where job_id = ?"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
		jobId,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.Task, 0)
	for rows.Next() {
		r := &database.Task{}
		if err := rows.Scan(&r.Id, &r.JobId, &r.Revision, &r.Success, &r.LogFile, &r.Command, &r.Target, &r.Via, &r.ConfigName, &r.StartAt, &r.FinishedAt, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}
	if len(res) > 0 {
		for _, v := range res {
			{
				rel, err := d.job.Select(ctx, v.JobId)
				if err != nil {
					return nil, xerrors.Errorf(": %w", err)
				}
				v.Job = rel
			}

		}
	}

	return res, nil

}

func (d *Task) ListPending(ctx context.Context, opt ...ListOption) ([]*database.Task, error) {
	listOpts := newListOpt(opt...)
	query := "select id, job_id, revision, success, log_file, command, target, via, config_name, start_at, finished_at, created_at, updated_at from task where start_at is null"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.Task, 0)
	for rows.Next() {
		r := &database.Task{}
		if err := rows.Scan(&r.Id, &r.JobId, &r.Revision, &r.Success, &r.LogFile, &r.Command, &r.Target, &r.Via, &r.ConfigName, &r.StartAt, &r.FinishedAt, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}
	if len(res) > 0 {
		for _, v := range res {
			{
				rel, err := d.job.Select(ctx, v.JobId)
				if err != nil {
					return nil, xerrors.Errorf(": %w", err)
				}
				v.Job = rel
			}

		}
	}

	return res, nil

}

func (d *Task) Create(ctx context.Context, task *database.Task, opt ...ExecOption) (*database.Task, error) {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(
		ctx,
		"INSERT INTO `task` (`job_id`, `revision`, `success`, `log_file`, `command`, `target`, `via`, `config_name`, `start_at`, `finished_at`, `created_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		task.JobId, task.Revision, task.Success, task.LogFile, task.Command, task.Target, task.Via, task.ConfigName, task.StartAt, task.FinishedAt, time.Now(),
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return nil, sql.ErrNoRows
	}

	task = task.Copy()
	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	task.Id = int32(insertedId)

	task.ResetMark()
	return task, nil

}

func (d *Task) Delete(ctx context.Context, id int32, opt ...ExecOption) error {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(ctx, "DELETE FROM `task` WHERE `id` = ?", id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (d *Task) Update(ctx context.Context, task *database.Task, opt ...ExecOption) error {
	if !task.IsChanged() {
		return nil
	}

	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	changedColumn := task.ChangedColumn()
	cols := make([]string, len(changedColumn)+1)
	values := make([]interface{}, len(changedColumn)+1)
	for i := range changedColumn {
		cols[i] = "`" + changedColumn[i].Name + "` = ?"
		values[i] = changedColumn[i].Value
	}
	cols[len(cols)-1] = "`updated_at` = ?"
	values[len(values)-1] = time.Now()

	query := fmt.Sprintf("UPDATE `task` SET %s WHERE `id` = ?", strings.Join(cols, ", "))
	res, err := conn.ExecContext(
		ctx,
		query,
		append(values, task.Id)...,
	)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	task.ResetMark()
	return nil

}

type TrustedUser struct {
	conn *sql.DB
}

type TrustedUserInterface interface {
	Tx(ctx context.Context, fn func(tx *sql.Tx) error) error
	Select(ctx context.Context, id int32) (*database.TrustedUser, error)
	ListAll(ctx context.Context, opt ...ListOption) ([]*database.TrustedUser, error)
	ListByGithubId(ctx context.Context, githubId int64, opt ...ListOption) ([]*database.TrustedUser, error)
	Create(ctx context.Context, trustedUser *database.TrustedUser, opt ...ExecOption) (*database.TrustedUser, error)
	Update(ctx context.Context, trustedUser *database.TrustedUser, opt ...ExecOption) error
	Delete(ctx context.Context, id int32, opt ...ExecOption) error
}

var _ TrustedUserInterface = &TrustedUser{}

func NewTrustedUser(conn *sql.DB) *TrustedUser {
	return &TrustedUser{
		conn: conn,
	}
}

func (d *TrustedUser) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if err := fn(tx); err != nil {
		rErr := tx.Rollback()
		return xerrors.Errorf("%v: %w", rErr, err)
	}

	err = tx.Commit()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil

}

func (d *TrustedUser) Select(ctx context.Context, id int32) (*database.TrustedUser, error) {
	row := d.conn.QueryRowContext(ctx, "SELECT * FROM `trusted_user` WHERE `id` = ?", id)

	v := &database.TrustedUser{}
	if err := row.Scan(&v.Id, &v.GithubId, &v.Username, &v.CreatedAt, &v.UpdatedAt); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	v.ResetMark()
	return v, nil

}

func (d *TrustedUser) ListAll(ctx context.Context, opt ...ListOption) ([]*database.TrustedUser, error) {
	listOpts := newListOpt(opt...)
	query := "select id, github_id, username, created_at, updated_at from trusted_user"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.TrustedUser, 0)
	for rows.Next() {
		r := &database.TrustedUser{}
		if err := rows.Scan(&r.Id, &r.GithubId, &r.Username, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}

	return res, nil

}

func (d *TrustedUser) ListByGithubId(ctx context.Context, githubId int64, opt ...ListOption) ([]*database.TrustedUser, error) {
	listOpts := newListOpt(opt...)
	query := "select id, github_id, username, created_at, updated_at from trusted_user where github_id = ?"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
		githubId,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.TrustedUser, 0)
	for rows.Next() {
		r := &database.TrustedUser{}
		if err := rows.Scan(&r.Id, &r.GithubId, &r.Username, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}

	return res, nil

}

func (d *TrustedUser) Create(ctx context.Context, trustedUser *database.TrustedUser, opt ...ExecOption) (*database.TrustedUser, error) {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(
		ctx,
		"INSERT INTO `trusted_user` (`github_id`, `username`, `created_at`) VALUES (?, ?, ?)",
		trustedUser.GithubId, trustedUser.Username, time.Now(),
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return nil, sql.ErrNoRows
	}

	trustedUser = trustedUser.Copy()
	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	trustedUser.Id = int32(insertedId)

	trustedUser.ResetMark()
	return trustedUser, nil

}

func (d *TrustedUser) Delete(ctx context.Context, id int32, opt ...ExecOption) error {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(ctx, "DELETE FROM `trusted_user` WHERE `id` = ?", id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (d *TrustedUser) Update(ctx context.Context, trustedUser *database.TrustedUser, opt ...ExecOption) error {
	if !trustedUser.IsChanged() {
		return nil
	}

	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	changedColumn := trustedUser.ChangedColumn()
	cols := make([]string, len(changedColumn)+1)
	values := make([]interface{}, len(changedColumn)+1)
	for i := range changedColumn {
		cols[i] = "`" + changedColumn[i].Name + "` = ?"
		values[i] = changedColumn[i].Value
	}
	cols[len(cols)-1] = "`updated_at` = ?"
	values[len(values)-1] = time.Now()

	query := fmt.Sprintf("UPDATE `trusted_user` SET %s WHERE `id` = ?", strings.Join(cols, ", "))
	res, err := conn.ExecContext(
		ctx,
		query,
		append(values, trustedUser.Id)...,
	)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	trustedUser.ResetMark()
	return nil

}

type PermitPullRequest struct {
	conn *sql.DB
}

type PermitPullRequestInterface interface {
	Tx(ctx context.Context, fn func(tx *sql.Tx) error) error
	Select(ctx context.Context, id int32) (*database.PermitPullRequest, error)
	ListByRepositoryAndNumber(ctx context.Context, repository string, number int32, opt ...ListOption) ([]*database.PermitPullRequest, error)
	Create(ctx context.Context, permitPullRequest *database.PermitPullRequest, opt ...ExecOption) (*database.PermitPullRequest, error)
	Update(ctx context.Context, permitPullRequest *database.PermitPullRequest, opt ...ExecOption) error
	Delete(ctx context.Context, id int32, opt ...ExecOption) error
}

var _ PermitPullRequestInterface = &PermitPullRequest{}

func NewPermitPullRequest(conn *sql.DB) *PermitPullRequest {
	return &PermitPullRequest{
		conn: conn,
	}
}

func (d *PermitPullRequest) Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if err := fn(tx); err != nil {
		rErr := tx.Rollback()
		return xerrors.Errorf("%v: %w", rErr, err)
	}

	err = tx.Commit()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil

}

func (d *PermitPullRequest) Select(ctx context.Context, id int32) (*database.PermitPullRequest, error) {
	row := d.conn.QueryRowContext(ctx, "SELECT * FROM `permit_pull_request` WHERE `id` = ?", id)

	v := &database.PermitPullRequest{}
	if err := row.Scan(&v.Id, &v.Repository, &v.Number, &v.CreatedAt, &v.UpdatedAt); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	v.ResetMark()
	return v, nil

}

func (d *PermitPullRequest) ListByRepositoryAndNumber(ctx context.Context, repository string, number int32, opt ...ListOption) ([]*database.PermitPullRequest, error) {
	listOpts := newListOpt(opt...)
	query := "select id, repository, number, created_at, updated_at from permit_pull_request where repository = ? and number = ?"
	if listOpts.limit > 0 {
		order := "ASC"
		if listOpts.desc {
			order = "DESC"
		}
		query = query + fmt.Sprintf(" ORDER BY `id` %s LIMIT %d", order, listOpts.limit)
	}
	rows, err := d.conn.QueryContext(
		ctx,
		query,
		repository,
		number,
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	res := make([]*database.PermitPullRequest, 0)
	for rows.Next() {
		r := &database.PermitPullRequest{}
		if err := rows.Scan(&r.Id, &r.Repository, &r.Number, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		r.ResetMark()
		res = append(res, r)
	}

	return res, nil

}

func (d *PermitPullRequest) Create(ctx context.Context, permitPullRequest *database.PermitPullRequest, opt ...ExecOption) (*database.PermitPullRequest, error) {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(
		ctx,
		"INSERT INTO `permit_pull_request` (`repository`, `number`, `created_at`) VALUES (?, ?, ?)",
		permitPullRequest.Repository, permitPullRequest.Number, time.Now(),
	)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return nil, sql.ErrNoRows
	}

	permitPullRequest = permitPullRequest.Copy()
	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	permitPullRequest.Id = int32(insertedId)

	permitPullRequest.ResetMark()
	return permitPullRequest, nil

}

func (d *PermitPullRequest) Delete(ctx context.Context, id int32, opt ...ExecOption) error {
	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	res, err := conn.ExecContext(ctx, "DELETE FROM `permit_pull_request` WHERE `id` = ?", id)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (d *PermitPullRequest) Update(ctx context.Context, permitPullRequest *database.PermitPullRequest, opt ...ExecOption) error {
	if !permitPullRequest.IsChanged() {
		return nil
	}

	execOpts := newExecOpt(opt...)
	var conn execConn
	if execOpts.tx != nil {
		conn = execOpts.tx
	} else {
		conn = d.conn
	}

	changedColumn := permitPullRequest.ChangedColumn()
	cols := make([]string, len(changedColumn)+1)
	values := make([]interface{}, len(changedColumn)+1)
	for i := range changedColumn {
		cols[i] = "`" + changedColumn[i].Name + "` = ?"
		values[i] = changedColumn[i].Value
	}
	cols[len(cols)-1] = "`updated_at` = ?"
	values[len(values)-1] = time.Now()

	query := fmt.Sprintf("UPDATE `permit_pull_request` SET %s WHERE `id` = ?", strings.Join(cols, ", "))
	res, err := conn.ExecContext(
		ctx,
		query,
		append(values, permitPullRequest.Id)...,
	)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if n, err := res.RowsAffected(); err != nil {
		return xerrors.Errorf(": %w", err)
	} else if n == 0 {
		return sql.ErrNoRows
	}

	permitPullRequest.ResetMark()
	return nil

}
