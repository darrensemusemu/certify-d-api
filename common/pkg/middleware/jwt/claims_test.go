package jwt

import (
	"context"
	"fmt"
	"os"
	"time"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/darrensemusemu/certify-d-api/common/pkg/env"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/matryer/is"
)

func TestClaimsGetJWKS(t *testing.T) {
	is := is.New(t)
	ts := testHelperAuthServer()
	defer ts.Close()

	claims := testHelperClaimsValidator(t, ts)
	keyFunc, err := claims.GetJWKS()
	is.NoErr(err)
	is.True(keyFunc != nil)
}

func TestClaimsParse(t *testing.T) {
	is := is.New(t)
	ts := testHelperAuthServer()
	defer ts.Close()

	tests := []struct {
		name         string
		jwtB64       string
		expectErr    bool
		overrideTime bool
		iss          string
	}{
		{
			name:         "jwt expired",
			jwtB64:       "eyJhbGciOiJSUzI1NiIsImtpZCI6IjdiY2UzNDhlLWJmNDQtNDU2Ni05OThkLTg5N2MxNmQ0NTRiNyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc0OTY5NzksImlhdCI6MTY1NzQ5MzM3OSwiaXNzIjoiaHR0cHM6Ly9jZXJ0aWZ5LWQuZGFycmVuc2VtdXNlbXUuZGV2IiwianRpIjoiMDBjNDYxZWQtZGY0NS00YmE2LTljMDQtMDBlNThlOWU2OGQ2IiwibmJmIjoxNjU3NDkzMzc5LCJzZXNzaW9uIjp7ImFjdGl2ZSI6dHJ1ZSwiYXV0aGVudGljYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNzE2OTg2WiIsImF1dGhlbnRpY2F0aW9uX21ldGhvZHMiOlt7ImFhbCI6ImFhbDEiLCJjb21wbGV0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY4NDI5MzU0MloiLCJtZXRob2QiOiJwYXNzd29yZCJ9XSwiYXV0aGVudGljYXRvcl9hc3N1cmFuY2VfbGV2ZWwiOiJhYWwxIiwiZXhwaXJlc19hdCI6IjIwMjItMDctMTFUMTU6Mzg6NDQuNjgzOFoiLCJpZCI6IjMxMmExMmQzLWQzNWUtNGQxMy1iMzA4LTA4YzRlM2JhNzczNiIsImlkZW50aXR5Ijp7ImNyZWF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY0NTc1N1oiLCJpZCI6IjI3MTE2MmJiLWEzMDgtNDQzMC04N2QwLWE3MWEzOWI4ZDY1ZiIsInJlY292ZXJ5X2FkZHJlc3NlcyI6W3siY3JlYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjYzMjk5WiIsImlkIjoiOGJiMThiOWQtZDM1Yi00OWNhLTgwYjctZGZmZDE2MDVlZDIzIiwidXBkYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjYzMjk5WiIsInZhbHVlIjoiZGFycmVuc2VtdXNlbXVAZ21haWwuY29tIiwidmlhIjoiZW1haWwifV0sInNjaGVtYV9pZCI6ImN1c3RvbWVyIiwic2NoZW1hX3VybCI6Imh0dHBzOi8va3JhdG9zLTY0NWY5NWQ1YmMtNnpnNGI6NDQzMy9zY2hlbWFzL1kzVnpkRzl0WlhJIiwic3RhdGUiOiJhY3RpdmUiLCJzdGF0ZV9jaGFuZ2VkX2F0IjoiMjAyMi0wNy0xMFQxNTozODo0NC42MzQ2ODJaIiwidHJhaXRzIjp7ImVtYWlsIjoiZGFycmVuc2VtdXNlbXVAZ21haWwuY29tIn0sInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY0NTc1N1oiLCJ2ZXJpZmlhYmxlX2FkZHJlc3NlcyI6W3siY3JlYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjU1NTQ0WiIsImlkIjoiN2Q1Y2NjZjMtZWVmZC00YmU0LTllMzctOTk5ZDU4YmZmMWYxIiwic3RhdHVzIjoic2VudCIsInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY1NTU0NFoiLCJ2YWx1ZSI6ImRhcnJlbnNlbXVzZW11QGdtYWlsLmNvbSIsInZlcmlmaWVkIjpmYWxzZSwidmlhIjoiZW1haWwifV19LCJpc3N1ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY4MzhaIn0sInN1YiI6IjI3MTE2MmJiLWEzMDgtNDQzMC04N2QwLWE3MWEzOWI4ZDY1ZiJ9.si_nbQTU8V48nonkj--DgQ8YCLNxWiwDff2qyxq8wbgpXSq234Jj2am4Geal0Z_B0btQFeZZGvKqUZelVyQIM5QkO6fuXjhBalVUz8Vf1INb_CMdMj_WuK4HJm89YRH6NmNg_FmxXO-JfPvjTtyjHGne5qtupagaO_B8ogZAAKhGpb2LkYqtkV6KmAk8WejlLv2Uf_wHeYFb4ACLwLsHHtocfPj2i5SxqdFmOh7pBVQtj_QtQPaeEn115gB64hU_dFtPQhYwNef0C5bMrb-WpU6pxWoZtIS9FFfhYQSp6modaH0IKE5xK5S-IG-Y3RP-ZyiDi3zN4URNFbqChcaVvA",
			expectErr:    true,
			overrideTime: false,
			iss:          "https://certify-d.darrensemusemu.dev",
		},
		{
			name:         "jwt valid",
			jwtB64:       "eyJhbGciOiJSUzI1NiIsImtpZCI6IjdiY2UzNDhlLWJmNDQtNDU2Ni05OThkLTg5N2MxNmQ0NTRiNyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc0OTY5NzksImlhdCI6MTY1NzQ5MzM3OSwiaXNzIjoiaHR0cHM6Ly9jZXJ0aWZ5LWQuZGFycmVuc2VtdXNlbXUuZGV2IiwianRpIjoiMDBjNDYxZWQtZGY0NS00YmE2LTljMDQtMDBlNThlOWU2OGQ2IiwibmJmIjoxNjU3NDkzMzc5LCJzZXNzaW9uIjp7ImFjdGl2ZSI6dHJ1ZSwiYXV0aGVudGljYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNzE2OTg2WiIsImF1dGhlbnRpY2F0aW9uX21ldGhvZHMiOlt7ImFhbCI6ImFhbDEiLCJjb21wbGV0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY4NDI5MzU0MloiLCJtZXRob2QiOiJwYXNzd29yZCJ9XSwiYXV0aGVudGljYXRvcl9hc3N1cmFuY2VfbGV2ZWwiOiJhYWwxIiwiZXhwaXJlc19hdCI6IjIwMjItMDctMTFUMTU6Mzg6NDQuNjgzOFoiLCJpZCI6IjMxMmExMmQzLWQzNWUtNGQxMy1iMzA4LTA4YzRlM2JhNzczNiIsImlkZW50aXR5Ijp7ImNyZWF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY0NTc1N1oiLCJpZCI6IjI3MTE2MmJiLWEzMDgtNDQzMC04N2QwLWE3MWEzOWI4ZDY1ZiIsInJlY292ZXJ5X2FkZHJlc3NlcyI6W3siY3JlYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjYzMjk5WiIsImlkIjoiOGJiMThiOWQtZDM1Yi00OWNhLTgwYjctZGZmZDE2MDVlZDIzIiwidXBkYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjYzMjk5WiIsInZhbHVlIjoiZGFycmVuc2VtdXNlbXVAZ21haWwuY29tIiwidmlhIjoiZW1haWwifV0sInNjaGVtYV9pZCI6ImN1c3RvbWVyIiwic2NoZW1hX3VybCI6Imh0dHBzOi8va3JhdG9zLTY0NWY5NWQ1YmMtNnpnNGI6NDQzMy9zY2hlbWFzL1kzVnpkRzl0WlhJIiwic3RhdGUiOiJhY3RpdmUiLCJzdGF0ZV9jaGFuZ2VkX2F0IjoiMjAyMi0wNy0xMFQxNTozODo0NC42MzQ2ODJaIiwidHJhaXRzIjp7ImVtYWlsIjoiZGFycmVuc2VtdXNlbXVAZ21haWwuY29tIn0sInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY0NTc1N1oiLCJ2ZXJpZmlhYmxlX2FkZHJlc3NlcyI6W3siY3JlYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjU1NTQ0WiIsImlkIjoiN2Q1Y2NjZjMtZWVmZC00YmU0LTllMzctOTk5ZDU4YmZmMWYxIiwic3RhdHVzIjoic2VudCIsInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY1NTU0NFoiLCJ2YWx1ZSI6ImRhcnJlbnNlbXVzZW11QGdtYWlsLmNvbSIsInZlcmlmaWVkIjpmYWxzZSwidmlhIjoiZW1haWwifV19LCJpc3N1ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY4MzhaIn0sInN1YiI6IjI3MTE2MmJiLWEzMDgtNDQzMC04N2QwLWE3MWEzOWI4ZDY1ZiJ9.si_nbQTU8V48nonkj--DgQ8YCLNxWiwDff2qyxq8wbgpXSq234Jj2am4Geal0Z_B0btQFeZZGvKqUZelVyQIM5QkO6fuXjhBalVUz8Vf1INb_CMdMj_WuK4HJm89YRH6NmNg_FmxXO-JfPvjTtyjHGne5qtupagaO_B8ogZAAKhGpb2LkYqtkV6KmAk8WejlLv2Uf_wHeYFb4ACLwLsHHtocfPj2i5SxqdFmOh7pBVQtj_QtQPaeEn115gB64hU_dFtPQhYwNef0C5bMrb-WpU6pxWoZtIS9FFfhYQSp6modaH0IKE5xK5S-IG-Y3RP-ZyiDi3zN4URNFbqChcaVvA",
			expectErr:    false,
			overrideTime: true,
			iss:          "https://certify-d.darrensemusemu.dev",
		},
	}

	claims := testHelperClaimsValidator(t, ts)
	keyFunc, _ := claims.GetJWKS()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.overrideTime {
				testHelperOverrideJWTTime()
				defer testHelperResetJWTTime()
			}

			err := claims.Parse(tt.jwtB64, keyFunc)
			is.True(tt.expectErr == (err != nil))
			if tt.expectErr {
				return
			}
			is.True(claims.VerifyIssuer(tt.iss, true))
		})
	}
}

func TestKratosAuthClaims(t *testing.T) {
	is := is.New(t)
	ts := testHelperAuthServer()
	defer ts.Close()
	testHelperClaimsValidator(t, ts)

	type args = struct {
		authHeader string
		env        string
	}

	tests := []struct {
		name               string
		expectedStatusCode int
		overrideTime       bool
		args               args
	}{
		{
			name:               "protected route no auth header",
			expectedStatusCode: http.StatusOK,
			overrideTime:       false,
			args: args{
				authHeader: "",
				env:        env.Production,
			},
		},
		{
			name:               "protected route invalid auth header",
			expectedStatusCode: http.StatusForbidden,
			overrideTime:       false,
			args: args{
				authHeader: "Bearer Not.Valid.JWT",
				env:        env.Production,
			},
		},
		{
			name:               "protected route valid auth header",
			expectedStatusCode: http.StatusOK,
			overrideTime:       true,
			args: args{
				authHeader: "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjdiY2UzNDhlLWJmNDQtNDU2Ni05OThkLTg5N2MxNmQ0NTRiNyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc0OTY5NzksImlhdCI6MTY1NzQ5MzM3OSwiaXNzIjoiaHR0cHM6Ly9jZXJ0aWZ5LWQuZGFycmVuc2VtdXNlbXUuZGV2IiwianRpIjoiMDBjNDYxZWQtZGY0NS00YmE2LTljMDQtMDBlNThlOWU2OGQ2IiwibmJmIjoxNjU3NDkzMzc5LCJzZXNzaW9uIjp7ImFjdGl2ZSI6dHJ1ZSwiYXV0aGVudGljYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNzE2OTg2WiIsImF1dGhlbnRpY2F0aW9uX21ldGhvZHMiOlt7ImFhbCI6ImFhbDEiLCJjb21wbGV0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY4NDI5MzU0MloiLCJtZXRob2QiOiJwYXNzd29yZCJ9XSwiYXV0aGVudGljYXRvcl9hc3N1cmFuY2VfbGV2ZWwiOiJhYWwxIiwiZXhwaXJlc19hdCI6IjIwMjItMDctMTFUMTU6Mzg6NDQuNjgzOFoiLCJpZCI6IjMxMmExMmQzLWQzNWUtNGQxMy1iMzA4LTA4YzRlM2JhNzczNiIsImlkZW50aXR5Ijp7ImNyZWF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY0NTc1N1oiLCJpZCI6IjI3MTE2MmJiLWEzMDgtNDQzMC04N2QwLWE3MWEzOWI4ZDY1ZiIsInJlY292ZXJ5X2FkZHJlc3NlcyI6W3siY3JlYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjYzMjk5WiIsImlkIjoiOGJiMThiOWQtZDM1Yi00OWNhLTgwYjctZGZmZDE2MDVlZDIzIiwidXBkYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjYzMjk5WiIsInZhbHVlIjoiZGFycmVuc2VtdXNlbXVAZ21haWwuY29tIiwidmlhIjoiZW1haWwifV0sInNjaGVtYV9pZCI6ImN1c3RvbWVyIiwic2NoZW1hX3VybCI6Imh0dHBzOi8va3JhdG9zLTY0NWY5NWQ1YmMtNnpnNGI6NDQzMy9zY2hlbWFzL1kzVnpkRzl0WlhJIiwic3RhdGUiOiJhY3RpdmUiLCJzdGF0ZV9jaGFuZ2VkX2F0IjoiMjAyMi0wNy0xMFQxNTozODo0NC42MzQ2ODJaIiwidHJhaXRzIjp7ImVtYWlsIjoiZGFycmVuc2VtdXNlbXVAZ21haWwuY29tIn0sInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY0NTc1N1oiLCJ2ZXJpZmlhYmxlX2FkZHJlc3NlcyI6W3siY3JlYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjU1NTQ0WiIsImlkIjoiN2Q1Y2NjZjMtZWVmZC00YmU0LTllMzctOTk5ZDU4YmZmMWYxIiwic3RhdHVzIjoic2VudCIsInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY1NTU0NFoiLCJ2YWx1ZSI6ImRhcnJlbnNlbXVzZW11QGdtYWlsLmNvbSIsInZlcmlmaWVkIjpmYWxzZSwidmlhIjoiZW1haWwifV19LCJpc3N1ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY4MzhaIn0sInN1YiI6IjI3MTE2MmJiLWEzMDgtNDQzMC04N2QwLWE3MWEzOWI4ZDY1ZiJ9.si_nbQTU8V48nonkj--DgQ8YCLNxWiwDff2qyxq8wbgpXSq234Jj2am4Geal0Z_B0btQFeZZGvKqUZelVyQIM5QkO6fuXjhBalVUz8Vf1INb_CMdMj_WuK4HJm89YRH6NmNg_FmxXO-JfPvjTtyjHGne5qtupagaO_B8ogZAAKhGpb2LkYqtkV6KmAk8WejlLv2Uf_wHeYFb4ACLwLsHHtocfPj2i5SxqdFmOh7pBVQtj_QtQPaeEn115gB64hU_dFtPQhYwNef0C5bMrb-WpU6pxWoZtIS9FFfhYQSp6modaH0IKE5xK5S-IG-Y3RP-ZyiDi3zN4URNFbqChcaVvA",
				env:        env.Production,
			},
		},
		{
			name:               "protected route dev env",
			expectedStatusCode: http.StatusOK,
			overrideTime:       false,
			args: args{
				authHeader: "",
				env:        env.Development,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(env.EnvVarKey, tt.args.env)

			if tt.overrideTime {
				testHelperOverrideJWTTime()
				defer testHelperResetJWTTime()
			}

			url := fmt.Sprintf("%v/auth/kratos", ts.URL)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			is.NoErr(err)

			if tt.args.authHeader != "" {
				req.Header.Set("Authorization", tt.args.authHeader)
			}
			client := http.Client{}
			resp, err := client.Do(req)
			is.NoErr(err)
			is.Equal(tt.expectedStatusCode, resp.StatusCode)
		})
	}
}

func TestKratosClaimsFromCtx(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	claims := NewKratosClaims()

	type args struct{ ctx context.Context }
	tests := []struct {
		name      string
		expectErr bool
		args      args
		ctx       context.Context
	}{
		{
			name:      "ctx without value",
			expectErr: true,
			args:      args{ctx: context.Background()},
		},
		{
			name:      "ctx with value",
			expectErr: false,
			args:      args{ctx: context.WithValue(ctx, KratosClaimsCxtKey, claims)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := KratosClaimsFromCtx(tt.args.ctx)
			is.True(tt.expectErr == (err != nil))
			if tt.expectErr {
				return
			}
			is.Equal(res, claims)
		})
	}
}

func testHelperClaimsValidator(t *testing.T, ts *httptest.Server) *KratosClaims {
	t.Setenv(EnvVarJWKSUrl, fmt.Sprintf("%v/.well-known/jwks.json", ts.URL))
	return NewKratosClaims()
}

func testHelperAuthServer() *httptest.Server {
	r := chi.NewRouter()

	// public routes
	r.Group(func(r chi.Router) {
		r.Get("/.well-known/jwks.json", func(w http.ResponseWriter, r *http.Request) {
			jwksString, err := os.ReadFile("./data/id_token.jwks.json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(200)
			w.Write(jwksString)
		})

	})

	// test protected routes
	r.Group(func(r chi.Router) {
		r.Use(KratosAuthMiddleware)

		r.Get("/auth/kratos", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	return httptest.NewServer(r)
}

func testHelperOverrideJWTTime() {
	tt := time.Unix(1657496900, 0)
	jwt.TimeFunc = func() time.Time {
		return tt
	}

}

func testHelperResetJWTTime() {
	jwt.TimeFunc = time.Now
}
