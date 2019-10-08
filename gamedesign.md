# Game Design

The following document describes in detail (in english) the game engine

## Game Algorithm

There are different phases

`start`, `turn`(s), and `end`

## Turn Algorithm

There are phases to the turn. They follow an indirect pattern, because
additional phases can stack or otherwise change the order

`sunrise`, `main`, `play`, `trigger`, `attack`, `defend`, `combat`, `sunset`

some phases have a functionality that changes the game state, labeled exec

### Sunrise

Active player must choose an element to create this turn
Any player may activate Play with a Spell, or Trigger
When done go to Main

exec
  - create element selected, active player draws a card

### Main

Active player may activate Play with all Card types
Any player may activate Play with a Spell, or Trigger
When done go to Attack

### Attack

Active player may declare a Being that is in their Present and Awake to attack: go to Defend
Active player may pass: go to Sunset
Any player may activate Play with a Spell, or Trigger

### Defend

Unactive player(s) may declare a Being that is in their Present and Awake to block: go to Combat
Unactive player(s) may declare no block: go to Combat
Any player may activate Play with a Spell, or Trigger

### Combat

Any player may activate Play with a Spell, or Trigger
When done go to Main
exec
  - If Being is blocking
    - Beings deal damage to each other
  - else
    - Being deals damage to opponent player

### Sunset

Any player may activate Play with a Spell, or Trigger
When done, go to Sunrise with active player = opponent player

## Play/Trigger

play and trigger create a state which stacks whatever state was previously active
the stacked state does exec or pass until this state has
the previous state is not destroyed, stopped, or restarted
Any player may activate Play with a Spell, or Trigger, stacking again


# Game Engine

Each game simulation has it's own goroutine, with an input `chan`. Player Agents submit requests on the input channel. Package `ai` implements a Player Agent that receives game data as a `JsonWriter`, and writes requests to the game input.

## Game Requests
- do nothing (errors, chat, data...)
- create a stack
  - do not call OnStop() for stacked state
- resolve a state
  - if there is a stack: write game state
  - else: follow turn order go to
