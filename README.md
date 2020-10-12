<img src="/assets/cover.png" alt="Gauguin - Generate opengraph images at runtime" />

*Gauguin** (pronounced `/ˈɡoʊɡæ̃/`) is an high performances Golang server that generates dynamic **opengraph** images at runtime.

# Configuration
**Gauguin** follows the configuration specified in `gauguin.yaml` file. Let's take the following file as an example:

```yaml
version: 0.0.1
routes:
  - path: /article/opengraph
    params:
      - title
      - author
      - imageurl
    size: 1200x630
    template: ./templates/article/opengraph.tmpl
  - path: /user/opengraph
    params:
      - title
      - username
      - imageurl
    size: 1200x630
    template: ./templates/user/opengraph.tmpl
```

with the above configuration, **Gauguin** will generate the following routes:

- `/article/opengraph`
  Query parameters:
    - `title`
    - `author`
    - `imageurl`

- `/user/opengraph`
  Query parameters:
    - `title`
    - `username`
    - `imageurl`

# Templates
As seen in the [configuration](#configuration) section, you can create a template for every route.

A template is basically a Golang `.tmpl` file, for example:

`./templates/article/opengraph.tmpl`
```tmpl
<!DOCTYPE html>
<html>
  <head>
    <link href="/public/templates/articles/opengraph.css" rel="stylesheet" />
  </head>
  <body style="background-image:url({{.imageurl}})">
    <h1>{{.title}}</h1>
    <p>{{.author}}</p>
  </body>
</html>
```

as you can see, at the moment all the CSS **must** be inline or external. I'm working hard on that.

# License
**Gauguin** is distributed under the [GPLv3 open source license](/LICENSE.md).