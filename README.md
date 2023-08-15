## Intro
- It's a simple REPL("read-eval-print loop"). You type in a command, and the program evaluates it and prints the result. You can then type in another command, and so on..
- Based on the commands entered we call [PokeApi's](https://pokeapi.co/) endpoints and prints result.

## Setup
- Clone repo
  ```bash
  git clone https://github.com/Kireeti-28/pokedexcli.git
  ```
- Build & Start
  ```bash
  go build && ./pokedexcli
  ```

## Getting Started
- From here on we will start interacting with cli via commands.
- **Avaliable Commands:**
   - `help`: Prints the help menu
   - `map`: List some location areas
   - `mapb`: List some location areas backward
   - `explore {location_area}`: Lists the pokemon in location areas
   - `catch {pokemon_name}`: Attempts to catch pokemon and add it to your pokedex
   - `inspect {pokemon_name}`: View information about caught pokemon
   - `pokedex`: Lists all your pokemons in your pokedex
   - `exit`: Turns off the Pokedex
