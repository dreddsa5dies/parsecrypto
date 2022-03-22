[![Go Report Card](https://goreportcard.com/badge/github.com/dreddsa5dies/parsecrypto)](https://goreportcard.com/report/github.com/dreddsa5dies/parsecrypto) ![License](https://img.shields.io/badge/License-GPL-blue.svg) 

## Тестовое задание
<details>
  <summary>Текст</summary>

### Парсинг Cryptorank:
- Ресурс: [Cryptorank](https://cryptorank.io/)
- Данные для парсинга: Теги нескольких валют (первых трех)
- Метод парсинга: Любой
- Метод хранения полученных результатов: Запись в гугл таблицы (api) по запуску
- Количество столбцов 3: Наименование, Теги, Timestamp.
- Время на выполнение: Решает исполнитель

### Парсинг CoinGecko
- Ресурс: [Coingecko](https://www.coingecko.com/)
- Данные для парсинга: Валюты, их стоимость относительно доллара
- Метод парсинга: Любой
- Метод хранения полученных результатов: Запись в гугл таблицы(api) по запуску. 
- Количество столбцов: Наименование, Цена, Timestamp. (Должно выводиться за один запрос 65! валют, вместе с ценами).
- Время на выполнение: Решает исполнитель

</details>


## License
This project is licensed under GPL license. Please read the [LICENSE](https:/github.com/dreddsa5dies/parsecrypto/tree/master/LICENSE.md) file.