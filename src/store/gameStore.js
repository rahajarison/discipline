import { defineStore } from 'pinia';
import { v4 as uuidv4 } from 'uuid';
import { ROUND_START, ROUND_WIN } from '../utils/hitContexts';
import { EVENT } from '../utils/types';
import { SYSTEM } from '../utils/categories';

const MAX_STACK_SIZE = 200;

export const useGameStore = defineStore('game', {
    state: () => ({
        rounds: [],
        activePlayer: 'P1',
        undoStack: [],
        redoStack: [],
    }),
    actions: {
        initializeStore() {
            this.startNewRound();
        },
        addAction(action) {
            const currentRound = this.rounds[this.rounds.length - 1];
            const newAction = {
                ...action,
                id: uuidv4(),
                player: this.activePlayer,
                hitContext: action.hitContext || "On hit",
                timestamp: new Date().toISOString(),
                gameTime: action.gameTime || null,
            };
            currentRound.actions.push(newAction);
            // this.rounds = [...this.rounds];
            this.resetRedoStack();
            console.log('Action added to round:', currentRound, newAction);
        },
        markRoundWinner() {
            const currentRound = this.rounds[this.rounds.length - 1];
            currentRound.actions.push({
                id: uuidv4(),
                category: SYSTEM,
                type: EVENT,
                player: this.activePlayer,
                hitContext: ROUND_WIN,
                timestamp: new Date().toISOString(),
            });

            // Ajouter un nouveau round si possible
            // TODO mettre la logique de si le win du round précédent était par le même joueur?
            if (this.rounds.length < 3) {
                console.log(`Player ${this.activePlayer} wins the round!`);
                this.startNewRound();
            } else {
                console.log("Maximum number of rounds has been reached")
            }
        },
        startNewRound() {
            this.resetRedoStack();
            this.rounds.push({ actions: [] });
            const currentRound = this.rounds[this.rounds.length - 1];
            currentRound.actions.push({
                id: uuidv4(),
                category: SYSTEM,
                type: EVENT,
                hitContext: ROUND_START,
                player: "N/A",
                timestamp: new Date().toISOString(),
            });
        },
        togglePlayer() {
            this.activePlayer = this.activePlayer === 'P1' ? 'P2' : 'P1';
            console.log("Change player" + this.activePlayer);
        },
        undo() {
            const currentRound = this.rounds[this.rounds.length - 1];
            if (currentRound.actions.length > 0) {
                const lastAction = currentRound.actions.pop();
                if (this.redoStack.length >= MAX_STACK_SIZE) {
                    this.redoStack.shift(); // If we have too much redos, we remove the oldest one
                }
                this.redoStack.push(JSON.parse(JSON.stringify(lastAction)));
                console.log('Undo:', lastAction);
            } else {
                console.log('No actions to undo.');
            }
        },
        redo() {
            if (this.redoStack.length > 0) {
                const lastRedo = this.redoStack.pop();

                // Restaurer l'action depuis redoStack dans la timeline
                this.rounds[this.rounds.length - 1].actions.push(JSON.parse(JSON.stringify(lastRedo)));

                console.log('Redo executed:', lastRedo);
            } else {
                console.log('No actions to redo.');
            }
        },
        resetRedoStack() {
            this.redoStack.splice(0, this.redoStack.length); // Reset reactive
            console.log('Redo stack reset.');
        },
    },
    getters: {
        flatActionList(state) {
            // let list = [];

            // for (let i = 0; i < state.rounds; i++) {
            //     console.log("In round", i)
            //     const round = gameStore.rounds[i];

            //     for (let action of round.actions) {
            //         console.log("Push action", action);
            //         list.push({
            //             ...action,
            //             round: i + 1,
            //         });
            //     }
            // }

            // return list;
            const result = [];
            for (let roundIndex = 0; roundIndex < state.rounds.length; roundIndex++) {
                const round = state.rounds[roundIndex];
                for (let actionIndex = 0; actionIndex < round.actions.length; actionIndex++) {
                    const action = round.actions[actionIndex];
                    result.push({
                        ...action, // Copie les propriétés de l'action
                        round: roundIndex + 1, // Ajoute la référence du round
                    });
                }
            }
            return result;
        },
        canRedo(state) {
            return state.redoStack.length > 0;
        },
        canUndo(state) {
            const currentRound = state.rounds[state.rounds.length - 1];
            return currentRound && currentRound.actions.length > 0;
        },
    },
});
