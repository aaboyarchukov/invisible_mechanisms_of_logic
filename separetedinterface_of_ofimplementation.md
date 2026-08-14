# Отделение интерфейса от реализации

Отделение интерфейса от реализации - базовый принцип для построения гибких, тестируемых и расширяемых систем. Когда мы не зависим от реализации нам становится намного легче внедрять новый функционал/имплементацию, а также чинить возникшие проблемы.

Также при проектировании таких контрактов - мы создаем единый источник правды, к которому можем обращаться.

## Решение

Реализуем контракт по сохранению данных в хранилище - данная реализация - сохранение в БД:

Снаяала опишем контракт:

```go
type Storage interface {
	Save(ctx context.Context, data string) error
	Retrieve(ctx context.Context, id int) (string, error)
}
```

Реализация:

```sql
-- name: SaveData :exec
INSERT INTO data_table (
	data
) SELECT sqlc.arg(data)::text;

-- name: RetrieveByID :one
SELECT
	data
FROM data_table
WHERE id = sqlc.arg(id)::integer;
```

```go
type Repository struct {
	db *pgx.Conn
	queries *sqlc.Queries
}

func GetQueries(ctx context.Context, queries *sqlc.Queries) *sqlc.Queries {
	if tx := TxFromContext(ctx); tx != nil {
		return queries.WithTx(tx)
	}
	return queries
}

func (r *Repository) Save(ctx context.Context, data string) error {
	return GetQueries(ctx, r.queries).SaveData(ctx, data)
}

func (r *Repository) Retrieve(ctx context.Context, id int) (string, error) {
	data, err := GetQueries(ctx, r.queries).RetrieveByID(ctx, id)
	if err != nil {
		return NullData, err
	}

	return data, nil
}
```
