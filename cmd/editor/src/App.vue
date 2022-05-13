<template>
  <el-container id="container">
    <el-header id="header">
      <el-button @click="run" type="success" :icon="runBtnIcon" size="medium" :disabled="runBtnDisabled">{{runBtnText}}</el-button>
    </el-header>
    <el-main id="main">
      <codemirror id="editor" ref="editor"
                  v-model="script"
                  :options="options"/>
      <el-input id="output" wrap="off" readonly
                v-model="output"
                type="textarea"
                placeholder="标准输出为空"/>
    </el-main>
  </el-container>
</template>

<script>
import { codemirror } from 'vue-codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/idea.css'
import 'codemirror/mode/javascript/javascript'
import 'codemirror/addon/edit/matchbrackets'
import 'codemirror/addon/edit/closebrackets'
import 'codemirror/addon/scroll/simplescrollbars.css'
import 'codemirror/addon/scroll/simplescrollbars'
import 'codemirror/addon/hint/show-hint.css'
import 'codemirror/addon/hint/show-hint'
import 'codemirror/addon/hint/javascript-hint'
import {runScript} from "@/common/api"
import {Message} from "element-ui"

export default {
  components: {
    codemirror
  },
  mounted() {
    this.editor.on('inputRead', () => {
      this.editor.showHint()
    })
  },
  computed: {
    editor() {
      return this.$refs.editor.codemirror
    }
  },
  data() {
    return {
      options: {
        tabSize: 4,
        indentUnit: 4,
        theme: 'idea', // 主题
        mode: 'text/javascript', // 语言
        autofocus: true,
        lineNumbers: true, // 显示行号
        matchBrackets: true, // 括号匹配检测
        autoCloseBrackets: true, // 括号自动闭合
        scrollbarStyle: 'simple', // 滚动条样式
        hintOptions: {
          completeSingle: false
        }
      },
      script: '',
      output: '',
      runBtnDisabled: false,
      runBtnIcon: 'el-icon-s-promotion',
      runBtnText: '点击运行'
    }
  },
  methods: {
    run() {
      this.output = '正在运行...'
      this.runBtnDisabled = true
      this.runBtnIcon = 'el-icon-loading'
      this.runBtnText = '正在运行'

      runScript({
        script: this.script
      }).then(res => {
        this.output = res.output
        this.runBtnDisabled = false
        this.runBtnIcon = 'el-icon-s-promotion'
        this.runBtnText = '点击运行'
      }).catch(err => {
        console.log(err)
        Message({
          showClose: true,
          message: err.message,
          type: 'error'
        })
        this.output = ''
        this.runBtnDisabled = false
        this.runBtnIcon = 'el-icon-s-promotion'
        this.runBtnText = '点击运行'
      })
    }
  }
}
</script>

<style>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html, body {
  height: 100%;
}

#container {
  height: 100%;
}

#header {
  margin-top: 20px;
  height: fit-content !important;
}

#main {
  height: 100%;
  display: flex;
}

#editor {
  flex: 7;
  height: 100%;
  border: 1px solid #dcdfe6;
  margin-right: 10px;
  overflow: hidden;
}

#output {
  height: 100%;
  border-radius: 0;
  resize: none;
}
</style>

<style>
.el-textarea {
  flex: 3;
  height: 100%;
  margin-left: 10px;

}

.CodeMirror {
  height: 100%;
}
</style>
