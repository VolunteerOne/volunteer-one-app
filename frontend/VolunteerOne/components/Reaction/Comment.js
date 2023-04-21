
import React, { useState }  from "react";
import { Block, Button, Card, Text, Input, theme} from "galio-framework";
import { StyleSheet, TextInput, View, TouchableOpacity, Pressable } from "react-native";
import Icon from "../Icon";

const Comment = ({commentCount}) => {
  const [show, setShow] = useState(false);
  const [comment, setComment] = useState("");
  const [showComment, setShowComment] = useState(false);
  const [count, setCount] = useState(commentCount);


  function sendValues(comment) {
    console.log(comment);
    setCount(count => count + 1);
    console.log(count);
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
            {count == 0 ? 
            <Text style={styles.titleText}>Comment</Text>
            :
            <Text style={styles.titleText}>{count} Comments</Text>  
            }
          </Block>
    </TouchableOpacity>

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
      <Pressable onPress={() => {sendValues(comment); setComment('')}}>
      <Icon
            family="MaterialIcons"
            size={15}
            name="send"
            color="#32325D"
          />
      </Pressable>
    </View>

    <View style={styles.close}>
      {/* <Pressable onPress={() => setShow(!show)}>
      <Icon
            family="MaterialIcons"
            size={15}
            name="close"
            color="#32325D"
          />
      </Pressable> */}
    </View>

    <View style={styles.viewComments}>
      <Pressable onPress={() => setShowComment(!showComment)}>
        {/* <Text>view comments</Text> */}
      </Pressable>
      { showComment ? 
      <View>
        <Text>{comment}</Text>
        </View>

          : null }
    </View>


    </View>
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
        marginTop: -60,
        zIndex : 1,
    },
    comment: {
      borderColor: "#32325D",
      backgroundColor: "#FFFFFF",
      borderRadius: 8,
      width: 340,
      marginLeft: -90,
      marginTop: -0,
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
      height: 15,
      marginTop: -17
    },
    viewComments: {
      position: 'absolute',
      marginLeft: 200,
      height: 15,
      marginTop: -17
    },
});

export default Comment;
    