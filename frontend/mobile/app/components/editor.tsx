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

const windowHeight = Dimensions.get('window').height;
const customFont = `${Inter}
  * {
      font-family: 'Inter';
  }`;
const initialContent = `<h1>Hello World</h1>`;
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
    const content = useEditorContent(editor, { type: 'html' });
    const [editable, setEditable] = useState(true);

    return (
        <View style={{flex: 1, alignItems: 'center', width: '100%', backgroundColor: 'white'}}>
          <View style={{paddingBottom: 20}}></View>
          <View style={{flexDirection: 'column', width: 380, alignItems: 'flex-end'}}><Button onPress={() => setEditable(!editable)} title={editable ? "Edit" : "Save"}/></View>
          {editable 
            ? (content && <WebView style={{width: 380}} scrollEnabled={false} source={{html: `
                            <style>
                              ${customFont}
                            </style>
                            <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=3.0, minimum-scale=1.0, user-scalable=yes">
                            ${content}`}}>
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