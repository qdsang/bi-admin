<template>
  <div class="vue-codemirror-wrap">
    <textarea />
  </div>
</template>

<script>
var CodeMirror = require('codemirror/lib/codemirror.js')
require('codemirror/lib/codemirror.css')
// import 'assets/override/codemirror_theme.css'
import 'codemirror/mode/sql/sql'
import 'codemirror/addon/hint/show-hint'
import 'codemirror/addon/hint/show-hint.css'
import 'codemirror/addon/edit/matchbrackets'
import 'codemirror/addon/display/placeholder'

export default {
  props: {
    value: {
      type: String,
      default: ''
    },
    options: {
      type: Object,
      default: function() {
        return {
          mode: 'text/javascript',
          lineNumbers: true,
          lineWrapping: true,
          undoDepth: 200
        }
      }
    },
    hints: {
      type: Array,
      default: function() {
        return [{ text: 'NIU' }]
      }
    }
  },
  data: function() {
    return {
      skipNextChangeEvent: false
    }
  },
  watch: {
    'value': function(newVal, oldVal) {
      var editorValue = this.editor.getValue()
      if (newVal !== editorValue) {
        this.skipNextChangeEvent = true
        var scrollInfo = this.editor.getScrollInfo()
        this.editor.setValue(newVal)
        this.editor.scrollTo(scrollInfo.left, scrollInfo.top)
      }
    },
    'options': function(newOptions, oldVal) {
      if (typeof newOptions === 'object') {
        for (var optionName in newOptions) {
          if (Object.prototype.hasOwnProperty.call(newOptions, optionName)) {
            this.editor.setOption(optionName, newOptions[optionName])
          }
        }
      }
    }
  },
  ready: function() {
    var _this = this
    this.editor = CodeMirror.fromTextArea(this.$el.querySelector('textarea'), this.options)
    this.editor.setValue(this.value)
    this.editor.on('change', function(cm, change) {
      if (_this.skipNextChangeEvent) {
        _this.skipNextChangeEvent = false
        return
      }
      _this.value = cm.getValue()
      if (!_this.$emit === false) {
        _this.$emit('change', cm.getValue())
      }
      _this.handleHints(change)
    })
  },
  mounted: function() {
    var _this = this
    this.editor = CodeMirror.fromTextArea(this.$el.querySelector('textarea'), this.options)
    this.editor.setValue(this.value)
    this.editor.on('change', function(cm, change) {
      if (_this.skipNextChangeEvent) {
        _this.skipNextChangeEvent = false
        return
      }
      if (!_this.$emit === false) {
        _this.$emit('change', cm.getValue())
        _this.$emit('input', cm.getValue())
      }
      _this.handleHints(change)
    })
  },
  beforeDestroy: function() {
    if (this.editor) {
      this.editor.toTextArea()
    }
  },
  methods: {
    handleChange(cm) {

    },
    handleHints(change) {
      if (change.origin === '+input' && change.text[0] !== ';' && change.text[0].trim() !== '' && change.text[1] !== '') {
        this.editor.showHint({
          completeSingle: false,
          tables: this.hints
        })
      }
    }
  }
}
</script>
<style>
  .CodeMirror-code {
    font-family: Menlo, Monaco, Consolas, "Courier New", monospace;
  }
</style>
