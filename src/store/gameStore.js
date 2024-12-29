import { defineStore } from 'pinia';
import { v4 as uuidv4 } from 'uuid';

export const useGameStore = defineStore('game', {
    state: () => ({
        rounds: [{ actions: [
            {
                "id": 'match-start',
                "type": "Round start",
                "player": "P1",
                "timestamp": new Date().toISOString(),
             }
        ] }],
        activePlayer: 'P1',
        undoStack: [],
        redoStack: [],
    }),
    actions: {
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
            this.redoStack = []; // Réinitialise Redo
            console.log('Action added to round:', currentRound, newAction);
        },
        markRoundWinner(player) {
            const currentRound = this.rounds[this.rounds.length - 1];
            currentRound.actions.push({
                id: uuidv4(),
                type: 'Round Win',
                player,
                hitContext: "N/A",
                timestamp: new Date().toISOString(),
            });

            // Ajouter un nouveau round si nécessaire
            // TODO mettre la logique de si le win du round précédent était par le même joueur?
            if (this.rounds.length < 3) {
                this.rounds.push({ actions: [] });
            }
            console.log(`Player ${player} wins the round!`);
        },
        togglePlayer() {
            this.activePlayer = this.activePlayer === 'P1' ? 'P2' : 'P1';
            console.log("Change player" + this.activePlayer);
        },
        undo() {
            const currentRound = this.rounds[this.rounds.length - 1];
            if (currentRound.actions.length > 0) {
                const lastAction = currentRound.actions.pop();
                this.undoStack.push(lastAction);
                console.log('Undo:', lastAction);
            }
        },
        redo() {
            if (this.undoStack.length > 0) {
                const lastUndo = this.undoStack.pop();
                this.redoStack.push(lastUndo);
                this.rounds[this.rounds.length - 1].actions.push(lastUndo);
                console.log('Redo:', lastUndo);
            }
        },
    },
});
