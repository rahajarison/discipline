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
import { prepareKeyMapForDisplay } from '../utils/keymap';

export default {
    setup() {
        const keyMapList = prepareKeyMapForDisplay();

        return {
            keyMapList,
        };
    },
};
</script>
