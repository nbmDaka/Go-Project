package main

import (
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

//func (app *application) library(w http.ResponseWriter, r *http.Request) {
//	app.render(w, r, "library.page.tmpl", &templateData{})
//}
//
//func (app *application) international(w http.ResponseWriter, r *http.Request) {
//	app.render(w, r, "international.page.tmpl", &templateData{})
//}
//
//func (app *application) events(w http.ResponseWriter, r *http.Request) {
//	app.render(w, r, "events.page.tmpl", &templateData{})
//}

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
	app.render(w, r, "food.page.tmpl", &templateData{
		FoodsData: s,
	})
	flash := app.session.PopString(r, "flash")
	app.render(w, r, "foods.page.tmpl", &templateData{
		Flash: flash,
	})
}

func (app *application) createFood(w http.ResponseWriter, r *http.Request) {
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
	app.session.Put(r, "flash", "Foods successfully created!")
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
