package migrations

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const (
	addTransactionXIDColumns = `
		ALTER TABLE relation_tuple_transaction
			ADD COLUMN xid xid8 NOT NULL DEFAULT (pg_current_xact_id()),
			ADD COLUMN snapshot pg_snapshot`

	addTupleXIDColumns = `
		ALTER TABLE relation_tuple
			ADD COLUMN created_xid xid8,
			ADD COLUMN deleted_xid xid8 NOT NULL DEFAULT ('9223372036854775807');`

	addNamespaceXIDColumns = `
		ALTER TABLE namespace_config
			ADD COLUMN created_xid xid8,
			ADD COLUMN deleted_xid xid8 NOT NULL DEFAULT ('9223372036854775807');`

	addTransactionDefault = `
		ALTER TABLE relation_tuple_transaction
			ALTER COLUMN snapshot SET DEFAULT (pg_current_snapshot());`

	addRelationTupleDefault = `
		ALTER TABLE relation_tuple
			ALTER COLUMN created_xid SET DEFAULT (pg_current_xact_id());`

	addNamepsaceDefault = `
		ALTER TABLE namespace_config
			ALTER COLUMN created_xid SET DEFAULT (pg_current_xact_id());`

	backfillTransactionXids = `
		UPDATE relation_tuple_transaction SET xid = id::text::xid8, snapshot = CONCAT(id, ':', id, ':')::pg_snapshot
			WHERE snapshot IS NULL;`

	backfillTupleXIDColumns = `
		UPDATE relation_tuple
			SET deleted_xid = deleted_transaction::text::xid8,
			created_xid = created_transaction::text::xid8
			WHERE created_xid IS NULL;`

	backfillNamespaceXIDColumns = `
		UPDATE namespace_config
			SET deleted_xid = deleted_transaction::text::xid8,
			created_xid = created_transaction::text::xid8
			WHERE created_xid IS NULL;`

	addTransactionSnapshotNotNull = `
		ALTER TABLE relation_tuple_transaction ALTER COLUMN snapshot SET NOT NULL;`

	addTupleCreatedNotNull = `
		ALTER TABLE relation_tuple ALTER COLUMN created_xid SET NOT NULL;`

	addNamespaceCreatedtNotNull = `
		ALTER TABLE namespace_config ALTER COLUMN created_xid SET NOT NULL;`
)

func init() {
	if err := DatabaseMigrations.Register("add-xid-columns", "add-ns-config-id",
		noNonatomicMigration,
		func(ctx context.Context, tx pgx.Tx) error {
			for _, stmt := range []string{
				addTransactionXIDColumns,
				addTupleXIDColumns,
				addNamespaceXIDColumns,
				addTransactionDefault,
				addRelationTupleDefault,
				addNamepsaceDefault,
				backfillTransactionXids,
				backfillTupleXIDColumns,
				backfillNamespaceXIDColumns,
				addTransactionSnapshotNotNull,
				addTupleCreatedNotNull,
				addNamespaceCreatedtNotNull,
			} {
				if _, err := tx.Exec(ctx, stmt); err != nil {
					return err
				}
			}

			return nil
		}); err != nil {
		panic("failed to register migration: " + err.Error())
	}
}
