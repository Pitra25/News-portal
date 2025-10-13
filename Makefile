mfd-xml:
  @mfd-generator xml -c "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -m ./docs/model/newsportal.mfd -n "news:news,categories,tags"

mfd-model:
  @mfd-generator model -m ./docs/model/vfs.mfd -p db -o ./db

mfd-repo:
  @mfd-generator repo -m ./docs/model/vfs.mfd -p db -o ./db
