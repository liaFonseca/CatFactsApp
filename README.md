# Cat Facts App

This is a simple Application mainly written in GO where you can request random facts about cats in several different languages.

This application has a GO API developed with Gin framework that provides two endpoints to:
- Request Cat Facts from the meowfacts API.
- Serve a web page built with HTML, Javascript and Bootstrap.

## How To Run:
1. **Clone the Repository:** Clone the repository to your local machine using the following command:
```bash
git clone https://github.com/liaFonseca/CatFactsApp
```

2. **Navigate to the Directory:** Go to the directory where the cloned repository lives:
```bash
cd CatFactsApp   
```
   
3. **Run the Code:** The following command installs the depencies, compiles the application's code and runs the main Go package.
```bash
go run main.go
```

## How To Use:
### Endpoints:

#### GET /view/catFacts
- Serves a web page where you can select the amount of facts you want to request and the language they are writen in. To issue the request press the Get Cat Facts button. 

**Default behaviour:** If no input is provided the application will return one cat fact written in english.
- The webpage can be accessed through a web browser:

```bash
http://localhost:8080/view/catFacts
```
---
#### GET /api/catFacts
- Returns a list of Cat Facts.
- **URL Parameters** (optional):
  - `count`: number of facts you want to retrieve. The maximum number is 10.
  - `lang`: language the facts are written in.
    
If no parameter is provided the API returns one cat fact in english.

Example:
       
```bash
http://localhost:8080/api/catFacts?count=3&lang=por
```
       
Response
       
```json
{
    "message": [
        "Muitas pessoas temem contrair uma doença protozoária, a toxoplasmose, de gatos. Esta doença também pode afetar o ser humano, mas mais seriamente, pode causar defeitos congênitos no nascimento. A toxoplasmose é uma doença comum, às vezes transmitida pelas fezes dos gatos. É causada com mais frequência pela ingestão de carne crua ou malpassada. Mulheres grávidas e pessoas com baixo sistema imunológico não devem tocar na caixa de areia do gato. Fora isso, não há razão para que essas pessoas evitem gatos.",
        "Paracetamol e chocolate são venenos para os gatos.",
        "Gatos, principalmente os mais velhos, podem ter câncer. Muitas vezes esta doença pode ser tratada com sucesso."
    ]
}
```
       

                            
## Available Languages
The supported languages are the ones specified in the meowfacts API documentation (https://github.com/wh-iterabb-it/meowfacts).

The following table was copied form their README.md file in 24/04/2023. 

Both the LANG Variable and ISO 639-1 values can be used.
  
| LANG Variable | [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) | Language Name | Localization (country) |
| :---: | :---: | :---: | :---: |
| `eng-us`  | `eng` | English | USA |
| `ces-cz` or `cze-cz` | `cze` (B) `ces` (T) | Czech | CZ |
| `ger-de`  | `ger` | German | DE |
| `ben-in`  | `ben` | Bengali | IN |
| `esp-es`  | `esp` | Spanish | ES |
| `esp-mx`  | `esp` | Spanish (default) | MX |
| `rus-ru`  | `rus` | Russian | RUS |
| `por-br`  | `por` | Portuguese | BR |
| `tl-fil`  | `fil` | Filipino | PH |
| `ukr-ua`  | `ukr` | Ukrainian | UA |
| `urd-ur`  | `urd` | Urdu | UR |
| `ita-it`  | `ita` | Italian | IT |
| `zho-tw` | `zho` | Chinese | TW (Taiwan) |
| `kor-ko`  | `kor` | Korean | KO |
