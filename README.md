# Pokedex CLI
A command-line interface (CLI) application written in Go that functions as a Pokedex. This application allows users to retrieve information about Pokemon using the PokeAPI (https://pokeapi.co/).

## Features 
- View areas containing Pokemon
- View, catch, and inspect Pokemon
- Cache frequently accessed Pokemon data for faster subsequent retrievals.

## Requirements
- Go (version 1.18 or later)
- Internet connection (to fetch data from the PokeAPI)
  
## Installation 
1. Clone the Repository:
```
git clone https://github.com/MechamJonathan/pokedexcli
cd pokedexcli
go build
```

## Run Application
In the pokedexcli directory run the following command:
```
./pokedexcli
```
# Commands
- help: Displays a help message
- catch <pokemon_name>: Attempt to catch a pokemon
- inspect <pokemon_name>: View details about a caught Pokemon
- pokedex: View caught Pokemon
- explore <location_name>: Explore a location
- map: Get the next page of locations
- mapb: Get the previous page of locations
- exit: Exit the Pokedex

## Example
```
Pokedex > map
  canalave-city-area
  eterna-city-area
  pastoria-city-area
  sunyshore-city-area
  sinnoh-pokemon-league-area
  oreburgh-mine-1f
  oreburgh-mine-b1f
  valley-windworks-area
  eterna-forest-area
  fuego-ironworks-area
  mt-coronet-1f-route-207
  mt-coronet-2f
  mt-coronet-3f
  mt-coronet-exterior-snowfall
  mt-coronet-exterior-blizzard
  mt-coronet-4f
  mt-coronet-4f-small-room
  mt-coronet-5f
  mt-coronet-6f
  mt-coronet-1f-from-exterior

Pokedex > explore mt-coronet-2f
  Exploring mt-coronet-2f...
  Found Pokemon: 
   - clefairy
   - golbat
   - machoke
   - graveler
   - nosepass
   - medicham
   - lunatone
   - solrock
   - chingling
   - bronzor
   - bronzong

Pokedex > catch lunatone
  Throwing a Pokeball at lunatone...
  lunatone escaped!

Pokedex > catch lunatone
  Throwing a Pokeball at lunatone...
  lunatone was caught!
  You may now inspect it with the inspect command

Pokedex > inspect lunatone
  Name: lunatone
  Height: 10
  Weight: 1680
  Stats:
    -hp: 90
    -attack: 55
    -defense: 65
    -special-attack: 95
    -special-defense: 85
    -speed: 70
  Types
    - rock
    - psychic

Pokedex > pokedex
  Your Pokedex:
   - pikachu
   - lunatone
```
