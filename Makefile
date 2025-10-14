mfd-xml:
	mfd-generator xml -c "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -m ./docs/model/newsportal.mfd -n "news:news,categories,tags"

mfd-model:
	mfd-generator model -m ./docs/model/newsportal.mfd -p db -o ./internal/db

mfd-repo:
	mfd-generator repo -m ./docs/model/newsportal.mfd -p db -o ./internal/db

mfd-test:
	mfd-generator dbtest -x News-portal/internal/db -o ./internal/db/test -m ./docs/model/newsportal.mfd