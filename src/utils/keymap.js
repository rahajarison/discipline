import { DEFENSE, NEUTRAL, SYSTEM } from "./categories";
import { ON_BLOCK, ON_HIT, WHIFF, ROUND_WIN, EVENT_NOTICEABLE } from "./hitContexts";
import { EVENT, REVIEWER } from "./types";

export const keyMap = {
    // Keys for managing actions
    a: {
        category: NEUTRAL,
        type: 'Jump In',
        hitContexts: {
            default: ON_HIT,
            ctrl: ON_BLOCK,
        },
    },
    z: {
        category: NEUTRAL,
        type: 'Dash In',
        hitContexts: {
            default: ON_HIT,
            ctrl: ON_BLOCK,
        },
    },
    e: {
        category: NEUTRAL,
        type: 'Normal confirm',
        hitContexts: {
            default: ON_HIT,
            ctrl: ON_BLOCK,
        },
    },
    r: {
        category: NEUTRAL,
        type: 'Skip Neutral',
        hitContexts: {
            default: ON_HIT,
            ctrl: ON_BLOCK,
        },
    },
    q: {
        category: NEUTRAL,
        type: 'Poke',
        hitContexts: {
            default: ON_HIT,
            ctrl: ON_BLOCK,
        },
    },
    s: {
        category: NEUTRAL,
        type: 'Cover Attack',
        hitContexts: {
            default: WHIFF,
        },
    },
    d: {
        category: DEFENSE,
        type: 'Throw Tech réussi',
        hitContexts: {
            default: ON_HIT,
        },
    },
    f: {
        category: NEUTRAL,
        type: 'Anti-Air réussi',
        hitContexts: {
            default: ON_HIT,
        },
    },
    g: {
        category: NEUTRAL,
        type: 'Punish',
        hitContexts: {
            default: ON_HIT,
        },
    },

    // Keys for moments
    Enter: {
        action: 'markRoundWinner',
    },
    " ": {
        type: REVIEWER,
        category: SYSTEM,
        hitContexts: {
            default: EVENT_NOTICEABLE,
        },
    },

    // Keys for management
    Tab: {
        action: 'togglePlayer',
    },
    ArrowDown: {
        action: 'redo',
    },
    ArrowUp: {
        action: 'undo',
    },
};


// Fonction pour déterminer le contexte
export const getHitContext = (event, hitContexts) => {
    console.log("Event modifiers:", {
        altKey: event.altKey,
        ctrlKey: event.ctrlKey,
    });
    console.log("Available hitContexts:", hitContexts);

    if (event.altKey && hitContexts.alt) {
        console.log("Returning alt hitContext:", hitContexts.alt);
        return hitContexts.alt;
    }
    if (event.ctrlKey && hitContexts.ctrl) {
        console.log("Returning ctrl hitContext:", hitContexts.ctrl);
        return hitContexts.ctrl;
    }
    if (hitContexts.default) {
        console.log("Returning default hitContext:", hitContexts.default);
        return hitContexts.default;
    }
    console.log("Fallback to ON_HIT");
    return ON_HIT; // Fallback
};

export const prepareKeyMapForDisplay = () => {
    const capitalizeFirstLetter = (str) => {
        if (!str) return '';
        return str.charAt(0).toUpperCase() + str.slice(1);
    };

    return Object.entries(keyMap).map(([key, details]) => {
        const modList = details.hitContexts
            ? Object.entries(details.hitContexts).map(([modifier, context]) => `${modifier}: ${context}`).join(', ')
            : 'N/A';

        return {
            key: capitalizeFirstLetter(key),         // Key (ex: "A")
            type: details.type || details.action,    // Type or action (ex: "Jump In")
            modifiers: modList,                      // Modifiers formatted (ex: "ctrl: On Block, default: On Hit")
        };
    });
};