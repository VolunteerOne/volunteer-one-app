
import React, { useState }  from "react";
import { Block, Button, Card, Text, Input, theme} from "galio-framework";
import { StyleSheet, TextInput, View, TouchableOpacity, Pressable } from "react-native";
import Icon from "../Icon";

const Comment = () => {
  const [show, setShow] = useState(false);
  const [comment, setComment] = useState("");

  function sendValues(comment) {
    console.log(comment);
};
  return (
      <View style={styles.view}>
      <TouchableOpacity
          style={styles.button}
          onPress={() => setShow(!show)}
          onChangeText={newText => setText(newText)}
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
                height: 40,
                marginLeft: 10
              }}
              onChangeText={newComment => setComment(newComment)}
              value={comment}
          />

    <View style={styles.submit}>
      <Pressable onPress={() => {sendValues(comment); setComment('') }}>
      <Icon
            family="MaterialIcons"
            size={15}
            name="send"
            color="#32325D"
          />
          

      </Pressable>
    </View>
{/* 
    <View style={styles.close}>
      <Pressable onPress={() => setShow(!show)}>
      <Icon
            family="MaterialIcons"
            size={15}
            name="close"
            color="#32325D"
          />

      </Pressable>
    </View> */}

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
      borderColor: "#32325D",
      backgroundColor: "#FFFFFF",
      borderRadius: 8,
      width: 340,
      marginLeft: -90,
      marginTop: -40,
      height: 40,
      zIndex: 5,
      flexDirection: 'row',
      borderColor: 'gray', borderWidth: 1

    },
    submit: {
      position: 'absolute',
      marginLeft: 315,
      height: 30,
      marginTop: 12
    },
    close: {
      position: 'absolute',
      marginLeft: 315,
      height: 50,
      marginTop: -17
    }
});

export default Comment;
    