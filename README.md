<img src="/assets/cover.png" alt="Gauguin - Generate opengraph images at runtime" />

<a href="https://www.producthunt.com/posts/gauguin?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-gauguin" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=276795&theme=dark" alt="Gauguin - Generate dynamic OpenGraph or social images in real-time | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54"  align="center"/></a>

**Gauguin** (pronounced `/ˈɡoʊɡæ̃/`) is an high performances Golang server that generates dynamic **opengraph** images at runtime.

🎉 Read the quickstart post [here](https://levelup.gitconnected.com/generate-dynamic-opengraph-images-using-gauguin-b53c5dc8ec2f)!

# Sponsors

<div align="center">
  <a href="https://graphcms.com?utm_source=https://github.com/micheleriva/krabs">
    <img src="/assets/sponsor-graphcms.svg" alt="GraphCMS" target="_blank" width="250px" />
  </a>

  <br />
  <a align="center" href="https://github.com/sponsors/micheleriva" target="_blank">
    <b>Become a sponsor</b>
  </a>
</div>

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

2) For each route, create a Golang `tmpl` file (named the same way you named it inside the configuration file).\
For example, this would be our `article.tmpl`:

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
3) Don't forget to create a file called `.env` and place there your [Env Variables](#Env-Variables). In this case, port will be 5419, as specified in [this](/docker-compose.yaml) docker-compose file that we are going to use.
4) Now it's time to run on terminal `docker-compose up -d`.
5) Choose a title, an author and an image for your article opengraph image. Pass them via querystring to the route you defined in your configuration file.
6) Go to `http://localhost:5491/articles/opengraph?author=Bojack%20Horseman&title=A%20Post%20About%20my%20Garden&imgUrl=https%3A%2F%2Fimages.unsplash.com%2Fphoto-1525498128493-380d1990a112%3Fixlib%3Drb-1.2.1%26ixid%3DeyJhcHBfaWQiOjEyMDd9%26auto%3Dformat%26fit%3Dcrop%26w%3D300%26q%3D80`
7) Admire the following image:

<img src="/assets/example.jpeg" alt="Gauguin opengraph image example" />

# Documentation

I'm currently writing more documentation, it will be available on **Gitbook**: [http://micheleriva.gitbook.io/gauguin](http://micheleriva.gitbook.io/gauguin)

# Env Variables
```env
PORT=<number>         # HTTPS Port
DOCKERIZED=<boolean>  # Set to true if running Gauguin in Docker
CHROME_URL=<string>   # Optional, the URL of the Chrome REST Debugging APIs 
```

# License
**Gauguin** is distributed under the [GPLv3 open source license](/LICENSE.md).
