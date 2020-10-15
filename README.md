<img src="/assets/cover.png" alt="Gauguin - Generate opengraph images at runtime" />

**Gauguin** (pronounced `/ˈɡoʊɡæ̃/`) is an high performances Golang server that generates dynamic **opengraph** images at runtime.

🎉 Read the quickstart post [here](https://www.hackdoor.io/articles/generate-dynamic-opengraph-images-using-gauguin-01d8592e58c0)!

# Gauguin in 6 easy steps

1) Create a configuration file called `gauguin.yaml`

```yml
version: 0.0.1
routes:
  - path: /articles/opengraph
    params:
      - title
      - author
      - imgUrl
    size: 1200x630
    template: ./templates/article.tmpl
  - path: /author/opengraph
    params:
      - username
      - imgUrl
    size: 1200x630
    template: ./templates/user.tmpl
```

2) For each route, create a Golang `tmpl` file (named the same way you named it inside the configuration file):

```html
<!DOCTYPE html>
<html>
  <head>
    <style>
      body {
        margin: 0;
        font-family: Arial;
        color: #fff;
      }
      .article-template {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        width: 1200px;
        height: 630px;
        background: #001f1c;
      }
      h1 {
        margin: 0;
        font-size: 32px;
      }
      img {
        width: 200px;
        height: 200px;
        object-fit: cover;
        border-radius: 15px;
        margin-bottom: 25px;
      }
    </style>
  </head>
  <body>
    <div class="article-template">
      <img src="{{.imgUrl}}" />
      <h1>{{.title}}</h1>
      <p>Written by <b>{{.author}}</b></p>
    </div>
  </body>
</html>
```

3) Copy [this](/docker-compose.yaml) docker-compose file locally and run `docker-compose up -d`
4) Choose a title, an author and an image for your article opengraph image. Pass them via querystrng to the route you defined in your configuration file.
5) Go to `http://localhost:5491/articles/test?author=Bojack%20Horseman&title=A%20Post%20About%20my%20Garden&imgUrl=https%3A%2F%2Fimages.unsplash.com%2Fphoto-1525498128493-380d1990a112%3Fixlib%3Drb-1.2.1%26ixid%3DeyJhcHBfaWQiOjEyMDd9%26auto%3Dformat%26fit%3Dcrop%26w%3D300%26q%3D80&dev=true`
6) Admire the following image:

<img src="/assets/example.jpg" alt="Gauguin opengraph image example" />

# Documentation

I'm currently writing more documentation, it will be available on **Gitbook**: [http://micheleriva.gitbook.io/gauguin](http://micheleriva.gitbook.io/gauguin)

# License
**Gauguin** is distributed under the [GPLv3 open source license](/LICENSE.md).
