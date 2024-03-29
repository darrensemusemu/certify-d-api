package jwt

import (
	"testing"

	"github.com/matryer/is"
)

func TestValidate(t *testing.T) {
	is := is.New(t)
	ts := testHelperAuthServer()
	defer ts.Close()

	testHelperOverrideJWTTime()
	defer testHelperResetJWTTime()

	jwtB64 := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjdiY2UzNDhlLWJmNDQtNDU2Ni05OThkLTg5N2MxNmQ0NTRiNyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc0OTY5NzksImlhdCI6MTY1NzQ5MzM3OSwiaXNzIjoiaHR0cHM6Ly9jZXJ0aWZ5LWQuZGFycmVuc2VtdXNlbXUuZGV2IiwianRpIjoiMDBjNDYxZWQtZGY0NS00YmE2LTljMDQtMDBlNThlOWU2OGQ2IiwibmJmIjoxNjU3NDkzMzc5LCJzZXNzaW9uIjp7ImFjdGl2ZSI6dHJ1ZSwiYXV0aGVudGljYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNzE2OTg2WiIsImF1dGhlbnRpY2F0aW9uX21ldGhvZHMiOlt7ImFhbCI6ImFhbDEiLCJjb21wbGV0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY4NDI5MzU0MloiLCJtZXRob2QiOiJwYXNzd29yZCJ9XSwiYXV0aGVudGljYXRvcl9hc3N1cmFuY2VfbGV2ZWwiOiJhYWwxIiwiZXhwaXJlc19hdCI6IjIwMjItMDctMTFUMTU6Mzg6NDQuNjgzOFoiLCJpZCI6IjMxMmExMmQzLWQzNWUtNGQxMy1iMzA4LTA4YzRlM2JhNzczNiIsImlkZW50aXR5Ijp7ImNyZWF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY0NTc1N1oiLCJpZCI6IjI3MTE2MmJiLWEzMDgtNDQzMC04N2QwLWE3MWEzOWI4ZDY1ZiIsInJlY292ZXJ5X2FkZHJlc3NlcyI6W3siY3JlYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjYzMjk5WiIsImlkIjoiOGJiMThiOWQtZDM1Yi00OWNhLTgwYjctZGZmZDE2MDVlZDIzIiwidXBkYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjYzMjk5WiIsInZhbHVlIjoiZGFycmVuc2VtdXNlbXVAZ21haWwuY29tIiwidmlhIjoiZW1haWwifV0sInNjaGVtYV9pZCI6ImN1c3RvbWVyIiwic2NoZW1hX3VybCI6Imh0dHBzOi8va3JhdG9zLTY0NWY5NWQ1YmMtNnpnNGI6NDQzMy9zY2hlbWFzL1kzVnpkRzl0WlhJIiwic3RhdGUiOiJhY3RpdmUiLCJzdGF0ZV9jaGFuZ2VkX2F0IjoiMjAyMi0wNy0xMFQxNTozODo0NC42MzQ2ODJaIiwidHJhaXRzIjp7ImVtYWlsIjoiZGFycmVuc2VtdXNlbXVAZ21haWwuY29tIn0sInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY0NTc1N1oiLCJ2ZXJpZmlhYmxlX2FkZHJlc3NlcyI6W3siY3JlYXRlZF9hdCI6IjIwMjItMDctMTBUMTU6Mzg6NDQuNjU1NTQ0WiIsImlkIjoiN2Q1Y2NjZjMtZWVmZC00YmU0LTllMzctOTk5ZDU4YmZmMWYxIiwic3RhdHVzIjoic2VudCIsInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY1NTU0NFoiLCJ2YWx1ZSI6ImRhcnJlbnNlbXVzZW11QGdtYWlsLmNvbSIsInZlcmlmaWVkIjpmYWxzZSwidmlhIjoiZW1haWwifV19LCJpc3N1ZWRfYXQiOiIyMDIyLTA3LTEwVDE1OjM4OjQ0LjY4MzhaIn0sInN1YiI6IjI3MTE2MmJiLWEzMDgtNDQzMC04N2QwLWE3MWEzOWI4ZDY1ZiJ9.si_nbQTU8V48nonkj--DgQ8YCLNxWiwDff2qyxq8wbgpXSq234Jj2am4Geal0Z_B0btQFeZZGvKqUZelVyQIM5QkO6fuXjhBalVUz8Vf1INb_CMdMj_WuK4HJm89YRH6NmNg_FmxXO-JfPvjTtyjHGne5qtupagaO_B8ogZAAKhGpb2LkYqtkV6KmAk8WejlLv2Uf_wHeYFb4ACLwLsHHtocfPj2i5SxqdFmOh7pBVQtj_QtQPaeEn115gB64hU_dFtPQhYwNef0C5bMrb-WpU6pxWoZtIS9FFfhYQSp6modaH0IKE5xK5S-IG-Y3RP-ZyiDi3zN4URNFbqChcaVvA"
	claims := testHelperClaimsValidator(t, ts)
	err := Validate(jwtB64, claims)
	is.NoErr(err)
}
