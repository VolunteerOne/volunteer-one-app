
import React, { useState }  from "react";
import { Block, Button, Card, Text, Input, theme} from "galio-framework";
import { StyleSheet, TextInput, View, TouchableOpacity } from "react-native";
import Icon from "../Icon";

const Comment = () => {
        const [show, setShow] = useState(false);
        const [text, setText] = useState('');


        return (
            <View style={styles.view}>
            <TouchableOpacity
                style={styles.button}
                onPress={() => setShow(!show)}
                onChangeText={newText => setText(newText)}
                defaultValue={text}
                >
                <Block flex row>
                <Icon
                  family="MaterialIcons"
                  size={15}
                  name="comment"
                  color="#32325D"
                />
                <Text style={styles.titleText}>Comment</Text>
                </Block>
          </TouchableOpacity>

          { show ? 
            <View style={styles.comment}>
                <TextInput
                    placeholder="Write a comment..."
                    style={{
                      alignSelf: 'baseline',
                      paddingBottom: -50,
                      height: 40
                        }}/>
                          <TouchableOpacity style={{position: 'relative', paddingTop: 10}} onPress={() => setShow(!show)}>
          <Text style={{ 
}}>                                                Close</Text>
      </TouchableOpacity>
            </View>

                : null }
</View>
        );
};
const styles = StyleSheet.create({
  view: {
    position: 'absolute',
    marginLeft: 75,

  },
    titleText: {
        fontSize: 15,
        color: "#32325D",
        textAlign: 'right',
        marginLeft: 7,
        marginTop: -1
      },
    button: {
        borderless: true,
        borderColor: '#fff',
        shadowColor: '#fff',
        backgroundColor: '#fff',
        height: 30,
        marginLeft: 120,
        marginTop: -30,
        zIndex : 1,
    },
    comment: {
      backgroundColor: "#FFFFFF",
      borderRadius: 8,
      width: 340,
      marginLeft: -90,
      marginTop: -40,
      height: 40,
      zIndex: 5,
      flexDirection: 'row'
    }

});

export default Comment;
    