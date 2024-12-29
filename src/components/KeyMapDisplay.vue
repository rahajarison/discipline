<template>
    <div class="keymap-container">
        <div class="keymap-display">
            <h2 class="keymap-title">Key Mapping</h2>
            <ul>
                <li v-for="keyItem in keyMapList" :key="keyItem.key">
                    <span class="key">{{ keyItem.key }}</span>:
                    <span class="type">{{ keyItem.type }}</span>,
                    <span class="modifiers">modificateurs: {{ keyItem.modifiers }}</span>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
import { prepareKeyMapForDisplay } from '../utils/keyMap';

export default {
    setup() {
        const keyMapList = prepareKeyMapForDisplay();

        return {
            keyMapList,
        };
    },
};
</script>

<style>
.keymap-display {
    margin: 1rem;
}

.keymap-display ul {
    list-style-type: none;
    padding: 0;
}

.keymap-display li {
    margin-bottom: 0.5rem;
    font-family: Arial, sans-serif;
}

.key {
    font-weight: bold;
    color: #007bff;
}

.type {
    font-style: italic;
    color: #28a745;
}

.modifiers {
    color: #6c757d;
}
</style>
