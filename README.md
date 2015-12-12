
# Использование языка go

Примеры приложений на языке [go](https://golang.org)

**Внимание**: Github не включает блоки кода в README.md. Полная версия этого файла доступна [тут](http://lekovr.github.io/golang-use/).

## Установка программной среды

Дистрибутив пока еще (версия 1.5.2) ставится из архива.

[Пример установки для 64bit Linux](https://gist.github.com/LeKovr/385074a9e60dbd179192#file-go-install-sh)

<script src="https://gist.github.com/LeKovr/385074a9e60dbd179192.js"></script>

Теперь для обновления версии go будет достаточно скачать и распаковать архив в /usr/local/go-VER и изменить ссылку /usr/local/go

## Приложение 1. Hello World

Исходный код первого приложения есть на странице загрузки дистрибутива.
Поместим его в [main.go](https://gist.github.com/LeKovr/869fb8adef79d2fc64eb#file-main-go):
<script src="https://gist.github.com/LeKovr/869fb8adef79d2fc64eb.js"></script>
и запустим
```
$ go run main.go
hello, world
```

## Приложение 2. Web

Для веб-приложения будем использовать фреймворк [Macaron](http://go-macaron.com/docs).
У него есть два [шаблонизатора](http://go-macaron.com/docs/middlewares/templating), мы выберем базовый - macaron.Renderer.

[main.go](https://gist.githubusercontent.com/LeKovr/74a8288a4cdcde0b1df3/raw/7b485f1a7dd59c39ae58a548fb513bb8965c7279/main.go):
<script src="https://gist.github.com/LeKovr/74a8288a4cdcde0b1df3/7b485f1a7dd59c39ae58a548fb513bb8965c7279.js"></script>

[templates/hello.tmpl](https://gist.github.com/LeKovr/85a8d17144b7c3bdc14a#file-hello-tmpl):
<script src="https://gist.github.com/LeKovr/85a8d17144b7c3bdc14a.js"></script>

Установим Macaron:
```
$ go get gopkg.in/macaron.v1
```
Запускаем
```
$ go run main.go
[Macaron] listening on 0.0.0.0:4000 (development)
```
Теперь можно открыть http://localhost:4000/

### Вишенка

Допустим, мы не хотим руками переключаться в браузер и копипастить туда адрес нашего сервиса.
Скачаем библиотеку -
```
$ go get github.com/skratchdot/open-golang/open
```
И добавим пару строк в код main.go, он станет [таким](https://gist.github.com/LeKovr/74a8288a4cdcde0b1df3#file-main-go):
<script src="https://gist.github.com/LeKovr/74a8288a4cdcde0b1df3.js"></script>

Запускаем
```
$ go run main.go
[Macaron] listening on 0.0.0.0:4000 (development)
[Macaron] Started GET / for 127.0.0.1
[Macaron] Completed / 200 OK in 844.828µs

```
Если нигде не ошиблись, браузер откроется сам.

## Copyright & License

Copyright (c) 2015 Alexey Kovrizhkin

License: [CC0 1.0 Universal](LICENSE)
