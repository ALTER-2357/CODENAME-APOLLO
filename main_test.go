package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	got := Ping()
	want := "PONG"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSDrecords(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/students/GET/?SID=1000", nil)
	response := httptest.NewRecorder()
	SDrecords(response, request)
	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "[\"{\\\"Achieved\\\":\\\"yes\\\",\\\"CID\\\":1,\\\"CNAME\\\":\\\"farming\\\",\\\"CSD\\\":\\\"01/09/2017\\\",\\\"Postcode\\\":\\\"bt40 3ua\\\",\\\"SID\\\":1000,\\\"SNAME\\\":\\\"kyle bell\\\",\\\"address\\\":\\\"5 ros na righ, mullagboy\\\",\\\"email\\\":\\\"kyle@kyle.com\\\",\\\"mobilenumber\\\":100100101010}\"]\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

}

func TestSDadduser(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "localhost:3031/students/adduser/?SID=1002SNAME=kyle b &address=5 ros na righ, mullagboy&Postcode=bt40 3ua&email=kyle@kyle.com&mobilenumber=100100101010&CNAME=farming&CID=000001&CSD=01/09/2017&achieved=yes", nil)
	response := httptest.NewRecorder()

	SDadduser(response, request)

	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "\"done\"\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		} else {
		}

	})

}

func TestSDdeleterow(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/students/DELETE/?SID=1002", nil)
	response := httptest.NewRecorder()

	SDadduser(response, request)

	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "\"done\"\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		} else {
		}

	})

}

func TestAadduser(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/AssessorDetails/adduser/?ANAME=KYLE&AID=100&address=1&Postcode=1&email=1&number=1&VID=1", nil)
	response := httptest.NewRecorder()

	Aadduser(response, request)

	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "\"done\"\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		} else {
		}

	})

}

func TestVDrecords(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/AssessorDetails/GET/?AID=100", nil)
	response := httptest.NewRecorder()

	Arecords(response, request)

	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "[\"{\\\"AID\\\":100,\\\"ANAME\\\":\\\"KYLE\\\",\\\"Postcode\\\":1,\\\"VID\\\":1,\\\"address\\\":1,\\\"email\\\":1,\\\"number\\\":1}\"]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		} else {
		}

	})

}

func TestVDrecords2(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/AssessorDetails/UPDATE/?ANAME=KYLE&AID=100&address=1&Postcode=1&email=1&number=1&VID=1", nil)
	response := httptest.NewRecorder()

	Arecords(response, request)

	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "\"done\"\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		} else {
		}

	})

}

func TestAdeleterow(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "localhost:3031/AssessorDetails/DELETE/?AID=100", nil)
	response := httptest.NewRecorder()

	Adeleterow(response, request)

	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "\"done\"\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		} else {
		}

	})

}

func TestCSDrecords(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/CourseScheduleDetails/GET/?CID=1", nil)
	response := httptest.NewRecorder()

	CSDrecords(response, request)
	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "[\"{\\\"CID\\\":1,\\\"CNAME\\\":\\\"KYLE\\\",\\\"CSD\\\":\\\"01/09/19\\\",\\\"cost\\\":0,\\\"duration\\\":\\\"3 years\\\"}\"]\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}

func TestCSDrecords2(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/CourseScheduleDetails/UPDATE/?CID=1&CNAME=KYLE&CSD=01/09/19&duration=3 years&cost=0", nil)
	response := httptest.NewRecorder()

	CSDrecords(response, request)
	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "\"done\"\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}

func TestCSDadduser(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/CourseScheduleDetails/adduser/?CID=2&CNAME=13&CSD=13&Postcode=13&duration=13&cost=13", nil)
	response := httptest.NewRecorder()

	Aadduser(response, request)

	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "\"done\"\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		} else {
		}

	})

}

func TestCSDdeleterow(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/CourseScheduleDetails/DELETE/?CID=122", nil)
	response := httptest.NewRecorder()

	Adeleterow(response, request)

	t.Run("testing status code", func(t *testing.T) {

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			t.Failed()
		}

	})
	t.Run("testing response", func(t *testing.T) {
		got := response.Body.String()
		want := "\"done\"\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		} else {
		}

	})
}
