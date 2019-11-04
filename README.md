# Database to markdown

Convert Database schema to Markdown (Only MySQL)

## Usage

> go run main.go root 123456 localhost database_name

## Requirements

- Go 1.13

## Example

| COLUMN_NAME | COLUMN_TYPE      | COLUMN_DEFAULT | IS_NULLABLE | COLUMN_KEY | EXTRA          | COLUMN_COMMENT |
|-------------|------------------|----------------|-------------|------------|----------------|----------------|
| table       | int(10) unsigned | 0              | YES         | PRI        | auto_increment | ID             |


## License

You can find the license for this code in [the LICENSE file](LICENSE).
