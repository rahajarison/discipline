// Table des touches et leurs contextes
export const keyMap = {
    a: {
        type: 'Jump In',
        hitContexts: {
            default: 'On Hit',
            ctrl: 'On Block',
        },
    },
    z: {
        type: 'Dash In',
        hitContexts: {
            default: 'On Hit',
            ctrl: 'On Block',
        },
    },
    e: {
        type: 'Normal Cancel par Drive Rush / Confirm',
        hitContexts: {
            default: 'On Hit',
            ctrl: 'On Block',
        },
    },
    r: {
        type: 'Skip Neutral',
        hitContexts: {
            default: 'On Hit',
            ctrl: 'On Block',
        },
    },
    q: {
        type: 'Poke',
        hitContexts: {
            default: 'On Hit',
            ctrl: 'On Block',
        },
    },
    s: {
        type: 'Cover Attack',
        hitContexts: {
            default: 'Whiff',
        },
    },
    d: {
        type: 'Throw Tech réussi',
        hitContexts: {
            default: 'On Hit',
        },
    },
    f: {
        type: 'Anti-Air réussi',
        hitContexts: {
            default: 'On Hit',
        },
    },
    g: {
        type: 'Punish',
        hitContexts: {
            default: 'On Hit',
        },
    },
    Enter: {
        type: 'Round Win',
    },
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
    console.log("Fallback to 'On Hit'");
    return 'On Hit'; // Fallback
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