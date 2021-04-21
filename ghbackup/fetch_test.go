package ghbackup

import (
	"net/http"
	"testing"
)

func Test_getNextURL(t *testing.T) {
  header := http.Header{}
  header.Add("Link", `<https://api.github.com/organizations/xxxxx/repos?per_page=100&page=2>; rel="next", <https://api.github.com/organizations/546750/repos?per_page=100&page=3>; rel="last"`)
  res := getNextURL(header)

  if res == "" {
		t.Error("Should get a second link")
  }
  if res != "https://api.github.com/organizations/xxxxx/repos?per_page=100&page=2" {
		t.Error("Should get the correct second link")
  }
}

func Test_getNextURLThirdPage(t *testing.T) {
  header := http.Header{}
  header.Add("Link", `<https://api.github.com/organizations/xxxxx/repos?per_page=100&page=1>; rel="prev", <https://api.github.com/organizations/xxxxx/repos?per_page=100&page=3>; rel="next", <https://api.github.com/organizations/xxxxx/repos?per_page=100&page=3>; rel="last", <https://api.github.com/organizations/xxxxx/repos?per_page=100&page=1>; rel="first"`)
  res := getNextURL(header)

  if res == "" {
		t.Error("Should get third link")
  }
  if res != "https://api.github.com/organizations/xxxxx/repos?per_page=100&page=3" {
		t.Error("Should get the correct third link", res)
  }
}
