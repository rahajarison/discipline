import { defineStore } from 'pinia';

export const useGameStore = defineStore('game', {
    state: () => ({
        rounds: [{ actions: [
            {
                "type": "Jump in",
                "player": "P1",
                "timestamp": "99"
             }
        ] }],
        activePlayer: 'P1',
    }),
    actions: {
        togglePlayer() {
            this.activePlayer = this.activePlayer === 'P1' ? 'P2' : 'P1';
            console.log("Change player" + this.activePlayer);
        },
        undo() {
            console.log('Undo action');
        },
        redo() {
            console.log('Redo action');
        },
    },
});
