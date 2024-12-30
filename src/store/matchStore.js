import { defineStore } from 'pinia';
import { v4 as uuidv4 } from 'uuid';


export const useMatchStore = defineStore('match', {
    state: () => ({
        matches: [],
        activeMatchId: null
    }),
    actions: {
        createMatch(replayId = null, url = null) {
            const newMatch = {
                id: uuidv4(),
                rounds: [],
                replayId,
                url,
                notes: '',
                reports: [],
                objectives: [],
            };
            this.matches.push(newMatch);
            this.activeMatchId = newMatch.id;
        },
        deleteMatch(matchId) {
            this.matches = this.matches.filter((match) => match.id !== matchId);
        },
        saveMatch(match) {
            if (!match.id) {
                match.id = match.id || uuidv4();
            }
            const existingIndex = this.matches.findIndex((m) => m.id === match.id);

            if (existingIndex !== -1) {
                this.matches[existingIndex] = { ...this.matches[existingIndex], ...match };
                console.log(`Match ${match.id} mis à jour.`);
            } else {
                this.matches.push(match);
                console.log(`Match ${match.id} créé.`);
            }
        }
    },
    getters: {
        activeMatch(state) {
            return state.matches.find((match) => match.id === state.activeMatchId);
        },
        activeRounds(state) {
            const match = state.matches.find((match) => match.id === state.activeMatchId);
            return match ? match.rounds : [];
        },
    },
});
