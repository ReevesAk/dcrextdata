// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// VSP is an object representing the database table.
type VSP struct {
	ID                   int              `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name                 string           `boil:"name" json:"name" toml:"name" yaml:"name"`
	APIEnabled           bool             `boil:"api_enabled" json:"api_enabled" toml:"api_enabled" yaml:"api_enabled"`
	APIVersionsSupported types.Int64Array `boil:"api_versions_supported" json:"api_versions_supported" toml:"api_versions_supported" yaml:"api_versions_supported"`
	Network              string           `boil:"network" json:"network" toml:"network" yaml:"network"`
	URL                  string           `boil:"url" json:"url" toml:"url" yaml:"url"`
	Launched             time.Time        `boil:"launched" json:"launched" toml:"launched" yaml:"launched"`

	R *vspR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L vspL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VSPColumns = struct {
	ID                   string
	Name                 string
	APIEnabled           string
	APIVersionsSupported string
	Network              string
	URL                  string
	Launched             string
}{
	ID:                   "id",
	Name:                 "name",
	APIEnabled:           "api_enabled",
	APIVersionsSupported: "api_versions_supported",
	Network:              "network",
	URL:                  "url",
	Launched:             "launched",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertypes_Int64Array struct{ field string }

func (w whereHelpertypes_Int64Array) EQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_Int64Array) NEQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_Int64Array) LT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Int64Array) LTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Int64Array) GT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Int64Array) GTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var VSPWhere = struct {
	ID                   whereHelperint
	Name                 whereHelperstring
	APIEnabled           whereHelperbool
	APIVersionsSupported whereHelpertypes_Int64Array
	Network              whereHelperstring
	URL                  whereHelperstring
	Launched             whereHelpertime_Time
}{
	ID:                   whereHelperint{field: `id`},
	Name:                 whereHelperstring{field: `name`},
	APIEnabled:           whereHelperbool{field: `api_enabled`},
	APIVersionsSupported: whereHelpertypes_Int64Array{field: `api_versions_supported`},
	Network:              whereHelperstring{field: `network`},
	URL:                  whereHelperstring{field: `url`},
	Launched:             whereHelpertime_Time{field: `launched`},
}

// VSPRels is where relationship names are stored.
var VSPRels = struct {
	VSPTicks string
}{
	VSPTicks: "VSPTicks",
}

// vspR is where relationships are stored.
type vspR struct {
	VSPTicks VSPTickSlice
}

// NewStruct creates a new relationship struct
func (*vspR) NewStruct() *vspR {
	return &vspR{}
}

// vspL is where Load methods for each relationship are stored.
type vspL struct{}

var (
	vspColumns               = []string{"id", "name", "api_enabled", "api_versions_supported", "network", "url", "launched"}
	vspColumnsWithoutDefault = []string{"name", "api_enabled", "api_versions_supported", "network", "url", "launched"}
	vspColumnsWithDefault    = []string{"id"}
	vspPrimaryKeyColumns     = []string{"id"}
)

type (
	// VSPSlice is an alias for a slice of pointers to VSP.
	// This should generally be used opposed to []VSP.
	VSPSlice []*VSP

	vspQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	vspType                 = reflect.TypeOf(&VSP{})
	vspMapping              = queries.MakeStructMapping(vspType)
	vspPrimaryKeyMapping, _ = queries.BindMapping(vspType, vspMapping, vspPrimaryKeyColumns)
	vspInsertCacheMut       sync.RWMutex
	vspInsertCache          = make(map[string]insertCache)
	vspUpdateCacheMut       sync.RWMutex
	vspUpdateCache          = make(map[string]updateCache)
	vspUpsertCacheMut       sync.RWMutex
	vspUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single vsp record from the query.
func (q vspQuery) One(ctx context.Context, exec boil.ContextExecutor) (*VSP, error) {
	o := &VSP{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for vsp")
	}

	return o, nil
}

// All returns all VSP records from the query.
func (q vspQuery) All(ctx context.Context, exec boil.ContextExecutor) (VSPSlice, error) {
	var o []*VSP

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VSP slice")
	}

	return o, nil
}

// Count returns the count of all VSP records in the query.
func (q vspQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count vsp rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q vspQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if vsp exists")
	}

	return count > 0, nil
}

// VSPTicks retrieves all the vsp_tick's VSPTicks with an executor.
func (o *VSP) VSPTicks(mods ...qm.QueryMod) vspTickQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"vsp_tick\".\"vsp_id\"=?", o.ID),
	)

	query := VSPTicks(queryMods...)
	queries.SetFrom(query.Query, "\"vsp_tick\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"vsp_tick\".*"})
	}

	return query
}

// LoadVSPTicks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (vspL) LoadVSPTicks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeVSP interface{}, mods queries.Applicator) error {
	var slice []*VSP
	var object *VSP

	if singular {
		object = maybeVSP.(*VSP)
	} else {
		slice = *maybeVSP.(*[]*VSP)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &vspR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &vspR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`vsp_tick`), qm.WhereIn(`vsp_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load vsp_tick")
	}

	var resultSlice []*VSPTick
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice vsp_tick")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on vsp_tick")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for vsp_tick")
	}

	if singular {
		object.R.VSPTicks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &vspTickR{}
			}
			foreign.R.VSP = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.VSPID {
				local.R.VSPTicks = append(local.R.VSPTicks, foreign)
				if foreign.R == nil {
					foreign.R = &vspTickR{}
				}
				foreign.R.VSP = local
				break
			}
		}
	}

	return nil
}

// AddVSPTicks adds the given related objects to the existing relationships
// of the vsp, optionally inserting them as new records.
// Appends related to o.R.VSPTicks.
// Sets related.R.VSP appropriately.
func (o *VSP) AddVSPTicks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*VSPTick) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.VSPID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"vsp_tick\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"vsp_id"}),
				strmangle.WhereClause("\"", "\"", 2, vspTickPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.VSPID = o.ID
		}
	}

	if o.R == nil {
		o.R = &vspR{
			VSPTicks: related,
		}
	} else {
		o.R.VSPTicks = append(o.R.VSPTicks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &vspTickR{
				VSP: o,
			}
		} else {
			rel.R.VSP = o
		}
	}
	return nil
}

// VSPS retrieves all the records using an executor.
func VSPS(mods ...qm.QueryMod) vspQuery {
	mods = append(mods, qm.From("\"vsp\""))
	return vspQuery{NewQuery(mods...)}
}

// FindVSP retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVSP(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*VSP, error) {
	vspObj := &VSP{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vsp\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, vspObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from vsp")
	}

	return vspObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *VSP) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no vsp provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(vspColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	vspInsertCacheMut.RLock()
	cache, cached := vspInsertCache[key]
	vspInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			vspColumns,
			vspColumnsWithDefault,
			vspColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(vspType, vspMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(vspType, vspMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vsp\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vsp\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into vsp")
	}

	if !cached {
		vspInsertCacheMut.Lock()
		vspInsertCache[key] = cache
		vspInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the VSP.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *VSP) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	vspUpdateCacheMut.RLock()
	cache, cached := vspUpdateCache[key]
	vspUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			vspColumns,
			vspPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update vsp, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vsp\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, vspPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(vspType, vspMapping, append(wl, vspPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update vsp row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for vsp")
	}

	if !cached {
		vspUpdateCacheMut.Lock()
		vspUpdateCache[key] = cache
		vspUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q vspQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for vsp")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for vsp")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VSPSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vspPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vsp\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, vspPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in vsp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all vsp")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *VSP) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no vsp provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(vspColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	vspUpsertCacheMut.RLock()
	cache, cached := vspUpsertCache[key]
	vspUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			vspColumns,
			vspColumnsWithDefault,
			vspColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			vspColumns,
			vspPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert vsp, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(vspPrimaryKeyColumns))
			copy(conflict, vspPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vsp\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(vspType, vspMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(vspType, vspMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert vsp")
	}

	if !cached {
		vspUpsertCacheMut.Lock()
		vspUpsertCache[key] = cache
		vspUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single VSP record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VSP) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no VSP provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), vspPrimaryKeyMapping)
	sql := "DELETE FROM \"vsp\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from vsp")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for vsp")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q vspQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no vspQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from vsp")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for vsp")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VSPSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no VSP slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vspPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vsp\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, vspPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from vsp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for vsp")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *VSP) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVSP(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VSPSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VSPSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vspPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vsp\".* FROM \"vsp\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, vspPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VSPSlice")
	}

	*o = slice

	return nil
}

// VSPExists checks if the VSP row exists.
func VSPExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vsp\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if vsp exists")
	}

	return exists, nil
}
