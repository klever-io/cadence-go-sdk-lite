module github.com/klever-io/cadence-go-sdk-lite

go 1.19

require (
	github.com/bits-and-blooms/bitset v1.2.2
	github.com/dave/dst v0.27.2
	github.com/fxamacker/cbor/v2 v2.4.1-0.20220515183430-ad2eae63303f
	github.com/go-test/deep v1.0.5
	github.com/leanovate/gopter v0.2.9
	github.com/logrusorgru/aurora v0.0.0-20200102142835-e9ef32dff381
	github.com/onflow/atree v0.4.0
	github.com/onflow/cadence v0.34.0
	github.com/rivo/uniseg v0.2.1-0.20211004051800-57c86be7915a
	github.com/stretchr/testify v1.7.3
	github.com/texttheater/golang-levenshtein/levenshtein v0.0.0-20200805054039-cae8b0eaed6c
	github.com/turbolent/prettier v0.0.0-20220320183459-661cc755135d
	go.opentelemetry.io/otel v1.8.0
	go.uber.org/goleak v1.1.10
	golang.org/x/crypto v0.1.0
	golang.org/x/exp v0.0.0-20221126150942-6ab00d035af9
	golang.org/x/text v0.4.0
	golang.org/x/tools v0.2.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fxamacker/circlehash v0.3.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.14 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/mod v0.6.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/onflow/cadence => github.com/klever-io/cadence-go-sdk-lite v1.0.0
	github.com/onflow/flow-go-sdk => github.com/klever-io/flow-go-sdk-lite v1.0.0

)
