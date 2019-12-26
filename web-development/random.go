func newTemplateCache() (map[string]*template.Template, map[int]string, error) {
  base_dir := "./ui/html/base"
  main_dir := "./ui/html"
  archive_dir := "./ui/html/archive"
  data_structures_dir := "./ui/html/archive/data-structures"
  algorithms_dir := "./ui/html/archive/algorithms"
  web_development_dir := "./ui/html/archive/web-development"
  concurrency_dir := "./ui/html/archive/concurrency"
  networking_dir := "./ui/html/archive/networking"
  all_dir := []string{main_dir, archive_dir, data_structures_dir, algorithms_dir, web_development_dir, concurrency_dir, networking_dir}

  cache := map[string]*template.Template{}
  page_index := make(map[int]string)
  index := 0
  for _, x := range all_dir {
    pages, err := filepath.Glob(filepath.Join(x, "*.page.tmpl"))
    if err != nil {
      return nil, nil, err
    }

    for _, page := range pages {
      name := filepath.Base(page)
      page_index[index] = name
      index++
      ts, err := template.ParseFiles(page)
      if err != nil {
        return nil, nil, err
      }

      ts, err = ts.ParseGlob(filepath.Join(base_dir, "*.layout.tmpl"))
      if err != nil {
        return nil, nil, err
      }

      ts, err = ts.ParseGlob(filepath.Join(base_dir, "*.partial.tmpl"))
      if err != nil {
        return nil, nil, err
      }

      cache[name] = ts
    }
  }

	return cache, page_index, nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
	buf.WriteTo(w)
}

func (app *application) random_post(w http.ResponseWriter, r *http.Request) {
	var page string
	limit := len(app.page_index)
	danger := []string{"learning.page.tmpl", "about.page.tmpl", "home.page.tmpl", "data-structures.page.tmpl", "algorithms.page.tmpl", "concurrency.page.tmpl", "web-development.page.tmpl", "networking.page.tmpl"}
	outloop:for{
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(limit)))
		if err != nil {
			app.serverError(w, err)
			return
		}
		n := nBig.Int64()
		page = app.page_index[int(n)]
		e := 0
		inloop:for _, v := range danger {
			if page != v {
				e++
				if e == len(danger) {
					break outloop
				}
			} else {
				break inloop
			}
		}
	}
	app.render(w, r, page)
}
