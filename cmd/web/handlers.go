package main

import (
	"AITUNews/pkg/forms"
	"AITUNews/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.news.First()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		NewsData: s,
	})
}

// createNews
func (app *application) createNews(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			app.serverError(w, err)
			return
		}
		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		imageUrl := r.PostForm.Get("image_url")

		id, err := app.news.Insert(title, content, imageUrl)
		if err != nil {
			app.serverError(w, err)
			return
		}

		app.session.Put(r, "flash", "News successfully created!")
		http.Redirect(w, r, fmt.Sprintf("/news/about?id=%d", id), http.StatusSeeOther)

	}
	app.render(w, r, "create.page.tmpl", &templateData{})
}

func (app *application) aboutPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.news.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	flash := app.session.PopString(r, "flash")

	app.render(w, r, "about.page.tmpl", &templateData{
		Flash: flash,
		News:  s,
	})

}

func (app *application) foods(w http.ResponseWriter, r *http.Request) {
	s, err := app.food.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	flash := app.session.PopString(r, "flash")
	app.render(w, r, "food.page.tmpl", &templateData{
		FoodsData: s,
		Flash:     flash,
	})
}

func (app *application) createFood(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			app.serverError(w, err)
			return
		}

		meal_name := r.PostForm.Get("meal_name")
		weekday := r.PostForm.Get("weekday")
		quantity := r.PostForm.Get("quantity")

		_, err = app.food.InsertFood(meal_name, weekday, quantity)
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/foods", http.StatusSeeOther)

		app.render(w, r, "createFood.page.tmpl", &templateData{})
	}
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password", "role")
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}

	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"), form.Get("role"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	s, err := app.users.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "authenticatedUserID", id)
	app.session.Put(r, "authenticatedUserRole", s.Role)
	if s.Role == "student" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/news/create", http.StatusSeeOther)
	}
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")

	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
