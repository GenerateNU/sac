import React, { useState } from "react";
import { KeyboardAvoidingView, Platform, View } from "react-native";
import {useEditorBridge, RichText, CoreBridge, useEditorContent, TenTapStartKit} from "@10play/tentap-editor";
import { StyleSheet } from "react-native";
import { Inter } from "./font";
import { Button } from "react-native";
import { WebView } from "react-native-webview";
import { CUSTOM_TOOLBAR } from "./tools";
import { Toolbar } from "./toolbar";
import {Dimensions} from 'react-native';
import {JsonContent,parseJsonToHtml } from "../../utils/jsonToHTML";

const json: JsonContent[] = [{"type":"doc","content":[{"type":"heading","attrs":{"level":1},"content":[{"type":"text","text":"Hello World"}]},{"type":"heading","attrs":{"level":2},"content":[{"type":"text","marks":[{"type":"bold"},{"type":"italic"},{"type":"strike"},{"type":"underline"}],"text":"This is generate"}]},{"type":"bulletList","content":[{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Task 1"}]}]},{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Task 2"}]}]}]},{"type":"paragraph"},{"type":"orderedList","attrs":{"start":1},"content":[{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Task 3"}]}]},{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Task 4"}]}]}]},{"type":"paragraph","content":[{"type":"text","marks":[{"type":"link","attrs":{"href":"https://generatenu.com","target":"_blank","rel":"noopener noreferrer nofollow","class":null}}],"text":"link"}]}]}]
const initialContent = parseJsonToHtml(json);

const windowHeight = Dimensions.get('window').height;
const customFont = `${Inter}
  * {
      font-family: 'Inter';
  }`;

  const offset = () => {
  return windowHeight > 700 ? 60 : 40;
}

export const Editor = () => {
    const editor = useEditorBridge({
      autofocus: true,
      avoidIosKeyboard: true,
      initialContent: initialContent,
      bridgeExtensions: [
        ...TenTapStartKit,
        CoreBridge.configureCSS(customFont),
      ],
    });
    
    const jsonOutput = useEditorContent(editor, { type: 'json' });
    const [editable, setEditable] = useState(true);
    const htmlOutput = useEditorContent(editor, { type: 'html' });
    
    return (
        <View style={{flex: 1, alignItems: 'center', width: '100%', backgroundColor: 'white'}}>
          <View style={{paddingBottom: 20}}></View>
          <View style={{flexDirection: 'column', width: 380, alignItems: 'flex-end'}}><Button onPress={() => setEditable(!editable)} title={editable ? "Edit" : "Save"}/></View>
          {editable 
            ? (jsonOutput && <WebView style={{width: 380}} scrollEnabled={false} source={{html: `
                            <style>
                              ${customFont}
                            </style>
                            <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=3.0, minimum-scale=1.0, user-scalable=yes">
                            ${htmlOutput}`}}>
                          </WebView>)
            : <>
                <RichText style={{width: 360}} editor={editor}/>
                <KeyboardAvoidingView
                  behavior={Platform.OS === 'ios' ? 'padding' : undefined}
                  style={styles.keyboardAvoidingView}
                  keyboardVerticalOffset={offset()}
                >
                    <Toolbar 
                      items={CUSTOM_TOOLBAR}
                      hidden={false} 
                      editor={editor} />
                </KeyboardAvoidingView>
              </>
          }
        </View>
    );
};

const styles = StyleSheet.create({
  keyboardAvoidingView: {
    position: 'absolute',
    width: '100%',
    bottom: 0,
  },
});