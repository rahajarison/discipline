import { defineStore } from 'pinia';

export const useGameStore = defineStore('game', {
    state: () => ({
        rounds: [{ actions: [
            {
                "id": 'unique-id',
                "type": "Jump in",
                "player": "P1",
                "timestamp": "99"
             }
        ] }],
        activePlayer: 'P1',
        undoStack: [],
        redoStack: [],
    }),
    actions: {
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
