package main

import "github.com/sashabaranov/go-openai"

const banner = "   ____ _           _    ____ ____ _____   ____  _          _ _ \n" +
	"  / ___| |__   __ _| |_ / ___|  _ \\_   _| / ___|| |__   ___| | |\n" +
	" | |   | '_ \\ / _` | __| |  _| |_) || |   \\___ \\| '_ \\ / _ \\ | |\n" +
	" | |___| | | | (_| | |_| |_| |  __/ | |    ___) | | | |  __/ | |\n" +
	"  \\____|_| |_|\\__,_|\\__|\\____|_|    |_|   |____/|_| |_|\\___|_|_|\n" +
	" *******************************GO******************************\n"

// client configuration constants
const model = openai.GPT3Dot5Turbo
const numberOfChoices = 1
const temperature = 0.7

const chatGptHeader = "ChatGpt:> %s\n"

// main shell constants
const mainShellWelcomeMessage = "\n\nWelcome to the coolest ChatGPT shell application in the world!\n" +
	"Digit 'help' to see the available commands\n" +
	"Digit 'quit' or 'exit' to quit the application\n"
const mainShellHeader = "\nshell:> "
const mainShellDefaultCaseMessage = "Command '%s' not implemented\n"
const mainShellHelpMessage = "---Available commands---\n" +
	"hello-world: Say hello to ChatGPT\n" +
	"echo: Activate echo mode\n" +
	"text-adventure: Play a text adventure\n"
const mainShellExitMessage = "Bye bye!\n"

// echo shell constants
const echoShellWelcomeMessage = "Echo mode activated: digit something to say to ChatGPT.\nDigit 'exit' or 'quit' to exit from echo mode.\n"
const echoShellHeader = "\necho:> "
const echoShellExitMessage = "Exiting echo mode\n"

// text-adventure shell constants
const textAdventureShellWelcomeMessage = "Welcome to the ChatGPT text adventure!\nInitializing ChatGPT...\n"
const textAdventureShellHeader = "\ntext-adventure:> "
const textAdventureInstructionsMessage = "Enter three keywords to generate your adventure\n"
const textAdventureExitMessage = "Exiting text-adventure\n"

// error messages
const genericErrorMessage = "An error occurred: %s\n"
const chatGptClientErrorMessage = "Error using openai API: %s\n"
const emptyStringMessage = "The message must not be empty: try another input\n"

// chat-gpt prompts
const helloWorldPrompt = "Hi ChatGPT, please, tell me Hello World in all spoken languages\n"
const taskPrompt = `Generate text adventure given requirements.
                        
            1. SETTING
            2. ALLOWED ACTIONS
            3. OBSTACLES
            4. POTENTIAL HARMFUL ACTIONS
            5. GOAL
                        
            Wait for requirements specification.`
const settingPrompt = `1. Setting is determined by three keywords inserted from user
                        
            user keywords:
            castle, woods, night
                        
            generated adventure: "You are driving your car at night, in a road in the middle of the woods. Suddenly, the car stops. Helped by the full moon, you see a secondary road the climb towards a castle."
                        
            Wait for my next requirement.`
const allowedActionPrompt = `2.You offer the player to write the action that he want to do.
                                    
            (start of example)

            What do you do?
             
            (end of example)
                        
            Wait for my next requirement.`
const obstaclesPrompt = `3.You must generate obstacles for the player.
            Obstacle can be riddles, enemies, hard choices.
                        
            (start of examples)
                        
            "the statue asks you 'what does weight more, a pound of feather or a pound of led?'"
                        
            "you face a pack of angry wolves"
                        
            "you must choose between saving the girl, or take the sword"
                        
            (end of examples)
                        
            Wait for my next requirement.`
const potentiallyHarmfulActionsPrompt = `4.Player is allowed to choose an action that is potentially harmful or fatal in the fictional context of the adventure.
            In case of fatal choice, is game over.
            In case of game over, you must prompt the user to try again the same adventure, or start a new one.
                        
            (start of example)
                        
            "You jump off the cliff, but the impact broke all your bones.
            Game over.
            Would you try again, or start a new adventure"
                        
            (end of example)
                        
            Wait for my next requirement.`
const goalPrompt = `5.The goal is to find a random rare artifact or defeat an enemy or do a particular quest.
            The artifact must be coherent with the adventure setting.
            The enemy must be a boss.
            The quest must be epic.
            You must not explicitly state the goal of the adventure.
            When the goal is reached, the player wins.`
const finalPrompt = `Ask the player for three keywords, generate the adventure, and prompt the player asking which action he wants to do`
