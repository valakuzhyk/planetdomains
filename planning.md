# Planet Domains
The aim is that this is a text game that multiple people can play.

## Client-Server model
Have a server with the game state, and allow interaction through rpc?
the rpc sends over the options that you have, and the client simply renders those options, allows you to pick them, sending back the response.

## Matchmaking
We can create rooms? Where people can play if they want.


## Designing the card interface

A card can 
* Either be a ship or a base
* Always has a play ability
* May have a ally ability
* May have a scrap ability


For the effects
  * I could create effect structs that each implement their own behavior
  * I could create a master list of abilities, and indicate which abilities the card specifically has.
    * I could use this same struct of effects as the current list of abilities that are available for you to choose.
    * Hard to implement or conditions



## Card Playing
Each turn, the game should autoplay the cards that don't need the user to make a decision.
 * Adding combat, trade or authority
 * Drawing a card
 * Causing the opponent to discard a card
 * 

What are other abilities?
 * Triggered on having enough bases
 * Triggered when buying something (put next ship or base on top of deck)
   * How will we tell where a card should be put after buying?
 * Destroy an enemies base
 * Copies base until the end of the turn
   * How do we handle the change of this to another base.
     * We need to ensure that the player knows what base this got transformed to,
     * Also the opponent needs to see the original base.
   * Potential solutions
     * Somehow change the card back at the end of the turn (more like reality)
     * Have a different view depending on which player's turn it is.
   * How do we trigger the end of turn switch?
     * Allow registering abilities on other players and evaluate abilities at the start of the turn.
   * These functions should have what parameters?
     * The field, the player, the card itself.
  
When you play a card, what happens with the ability?
  You should let the player activate abilities by specifying the card. In the future, you could also allow automatic activation if there is no reason not to activate them.