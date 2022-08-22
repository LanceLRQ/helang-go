package helang

import (
	assertProvider "github.com/stretchr/testify/assert"
	"helang-go/helang/core"
	"log"
	"testing"
)

func TestU8ParseDef(t *testing.T) {
	assert := assertProvider.New(t)

	code := `
        u8 list1 = 1 | 2 | 3;
        u8 list2 = [3];
	`
	env, _, err := RunCodeWithoutEnv(code)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(env["list1"], core.NewU8Array([]int{1, 2, 3}))
}

func TestParseU8Set(t *testing.T) {
	assert := assertProvider.New(t)

	code := `
        u8 a = 1 | 2 | 3;
        u8 b = 4 | 5 | 6;
        a[1 | 3] = 12;
        b[0] = 10;
	`
	env, _, err := RunCodeWithoutEnv(code)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(env["a"], core.NewU8Array([]int{12, 2, 12}))
	assert.Equal(env["b"], core.NewU8Array([]int{10, 10, 10}))
}

func TestParseU8Get(t *testing.T) {
	env := map[string]*core.U8 {
		"a": core.NewU8Array([]int {2, 3, 4}),
	}
	assert := assertProvider.New(t)

	code := `
        u8 b = a[1 | 3];
	`
	_, err := RunCode(code, env)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(env["b"], core.NewU8Array([]int{2, 4}))
}

func TestParseDecl(t *testing.T) {
	assert := assertProvider.New(t)
	env, _, err := RunCodeWithoutEnv("u8 a;")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(env["a"], core.NewU8Empty())
}

func TestParseVarAssign(t *testing.T) {
	assert := assertProvider.New(t)
	env, _, err := RunCodeWithoutEnv(`
		u8 a = 1 | 2;
        a = 3 | 4;
	`)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(env["a"], core.NewU8Array([]int{3, 4}))
}

func TestParseIncrement(t *testing.T) {
	assert := assertProvider.New(t)
	env, _, err := RunCodeWithoutEnv(`
		u8 a = 1 | 2 | 3;
        a++;
	`)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(env["a"], core.NewU8Array([]int{2, 3, 4}))
}

func TestParseSemicolon(t *testing.T) {
	assert := assertProvider.New(t)
	env, _, err := RunCodeWithoutEnv(";;;u8 a = 1;;;")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(env["a"], core.NewU8Array([]int{1}))
}

//func TestParseSub(t *testing.T) {
//	assert := assertProvider.New(t)
//	env, _, err := helang.RunCodeWithoutEnv(`
//        u8 a = 4 | 6 | 7;
//        u8 b = 2 | 3 | 4;
//        u8 c = 1;
//        u8 a_b = a - b;
//        u8 a_b_c = a - b - c;
//	`)
//	if err != nil {
//		log.Fatal(err)
//	}
//	assert.Equal(env["a_b"], core.NewU8Array([]int{2, 3, 3}))
//	assert.Equal(env["a_b_c"], core.NewU8Array([]int{1, 2, 2}))
//}

