<template>
  <div class="my-2">
    <button
      @click="toggleDropdown"
      class="inline-flex items-center text-gray-500 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-3 py-1.5 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
      type="button"
    >
      <!-- <svg
        class="w-3 h-3 text-gray-500 dark:text-gray-400 me-3"
        aria-hidden="true"
        xmlns="http://www.w3.org/2000/svg"
        fill="currentColor"
        viewBox="0 0 20 20"
      >
        <path
          d="M10 0a10 10 0 1 0 10 10A10.011 10.011 0 0 0 10 0Zm3.982 13.982a1 1 0 0 1-1.414 0l-3.274-3.274A1.012 1.012 0 0 1 9 10V6a1 1 0 0 1 2 0v3.586l2.982 2.982a1 1 0 0 1 0 1.414Z"
        />
      </svg> -->
      {{ selectedLabel || placeholder }}
      <svg
        class="w-2.5 h-2.5 ms-2.5"
        aria-hidden="true"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 10 6"
      >
        <path
          stroke="currentColor"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="m1 1 4 4 4-4"
        />
      </svg>
    </button>

    <!-- Dropdown -->
    <div
      v-if="isOpen"
      class="z-10 w-48 bg-white divide-y divide-gray-100 rounded-lg shadow dark:bg-gray-700 dark:divide-gray-600"
    >
      <ul class="p-3 space-y-1 text-sm text-gray-700 dark:text-gray-200">
        <li
          v-for="option in options"
          :key="option.value"
          class="flex items-center p-2 rounded hover:bg-gray-100 dark:hover:bg-gray-600"
        >
          <input
            type="radio"
            :id="option.value"
            :value="option.value"
            name="filter"
            class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
            v-model="selected"
          />
          <label
            :for="option.value"
            class="w-full ms-2 text-sm font-medium text-gray-900 rounded dark:text-gray-300"
          >
            {{ option.label }}
          </label>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    options: {
      type: Array,
      required: true,
    },
    placeholder: {
      type: String,
      default: "Selectionner...",
    },
    modelValue: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      isOpen: false,
      selected: this.modelValue,
    };
  },
  computed: {
    selectedLabel() {
      const selectedOption = this.options.find(
        (option) => option.value === this.selected
      );
      return selectedOption ? selectedOption.label : null;
    },
  },
  watch: {
    selected(newValue) {
      this.$emit('update:modelValue', newValue);
    },
  },
  methods: {
    toggleDropdown() {
      this.isOpen = !this.isOpen;
    },
  },
};
</script>
