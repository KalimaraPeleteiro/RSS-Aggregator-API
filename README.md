<h1 align="center"> RSS Aggregator-API </h1>

Uma API feita em Go com o pacote [Chi](https://github.com/go-chi/chi) que permite usuários seguirem novos feeds RSS, e, 
eventualmente, receber os últimos posts dos feeds que eles seguem. Originalmente, o projeto foi feito por [Lane Wagner](https://github.com/wagslane) em um [vídeo](https://www.youtube.com/watch?v=un6ZyFkqFKo) da freeCodeCamp Academy, mas eu adaptei o design original, adicionando novas rotas e modificando o comportamento de outras para a criação de uma UI.
<br>

> Para os que não conhecem, feeds RSS, ou *Really Simple Sindication* é uma forma de disponibilizar seu conteúdo em um formato de fácil acesso e interatividade para
> outras aplicações, como essa API. O conteúdo é disponibilizado em formato XML, fazendo uso de tags padronizadas para facilitar a coleta de conteúdo. Abaixo, temos
> um exemplo de reportagem, representada no RSS Feed da CNN, em XML.
> <br>
> ```
> <item>
>  <title>
> <![CDATA[ Analysis: Fox News is about to enter the true No Spin Zone ]]>
> </title>
> <description>
> <![CDATA[ This is it. ]]>
> </description>
> <link>https://www.cnn.com/2023/04/14/media/fox-news-dominion-hnk-intl/index.html</link>
> <guid isPermaLink="true">https://www.cnn.com/2023/04/14/media/fox-news-dominion-hnk-intl/index.html</guid>
> <pubDate>Fri, 14 Apr 2023 06:41:00 GMT</pubDate>
> <media:group>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-super-169.jpg" height="619" width="1100" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-large-11.jpg" height="300" width="300" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-vertical-large-gallery.jpg" height="552" width="414" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-video-synd-2.jpg" height="480" width="640" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-live-video.jpg" height="324" width="576" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-t1-main.jpg" height="250" width="250" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-vertical-gallery.jpg" height="360" width="270" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-story-body.jpg" height="169" width="300" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-t1-main.jpg" height="250" width="250" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-assign.jpg" height="186" width="248" type="image/jpeg"/>
> <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230413121210-01-dominion-courtroom-hp-video.jpg" height="144" width="256" type="image/jpeg"/>
> </media:group>
> </item>
> <item>
> ```

<br>
<br>
<h3>Principais Rotas</h3>

Para usuários, as rotas mais importantes são as que seguem abaixo. **Todas as rotas exigem um header de `Authorization` com sua chave específica de API, em um formato
`APIKey suachave`.**

> **Seguir um novo feed**
> <br>
> Caminho: *v1/users/follow*
> <br>
> É necessário passar o id do feed a ser seguido em formato JSON.
> ```
> {
>      "feed_id": "df173a4e-9c6e-4127-ad6d-05e55c99f957"
> }
> ```


> **Retornar todos os feeds que um usuário segue**
> <br>
> Caminho: *v1/users/my_feeds*


> **Deixar de seguir um feed**
> <br>
> Caminho: *v1/users/unfollow/id_do_feed*
> <br>
> É necessário passar o id do feed a deixar de ser seguido na rota. Segue um exemplo abaixo.
> ```
> v1/users/unfollow/d5400ea0-fbe7-4731-b0e9-d828b0ab7989
> ```


> **Retornar os posts dos feeds seguidos por um usuário**
> <br>
> Caminho: *v1/users/my_feeds/posts*




