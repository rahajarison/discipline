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
                timestamp: new Date(),
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
        },
        getFlatActionList(match) {
            const result = [];
            const matchInstance = this.findMatch(match.id);
            if (matchInstance) {
                for (let roundIndex = 0; roundIndex < matchInstance.rounds.length; roundIndex++) {
                    const round = matchInstance.rounds[roundIndex];
                    for (let actionIndex = 0; actionIndex < round.actions.length; actionIndex++) {
                        const action = round.actions[actionIndex];
                        result.push({
                            ...action,
                            round: roundIndex + 1,
                        });
                    }
                }
            }
            return result;
        },
        findMatch(matchId) {
            return this.matches.find((match) => match.id === matchId);
        }
    },
    getters: {
        activeMatch(state) {
            return state.findMatch(state.activeMatchId);
        },
        activeRounds(state) {
            const match = state.activeMatch;
            return match ? match.rounds : [];
        },
        flatActiveActionList(state) {
            return state.getFlatActionList(state.activeMatch);
        },
    },
});
