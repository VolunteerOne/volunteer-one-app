
import React, { useState }  from "react";
import { Block, Button, Card, Text, Input, theme} from "galio-framework";
import { StyleSheet, View, TextInput, TouchableOpacity, Pressable, Dimensions } from "react-native";
import Icon from "../Icon";
import { argonTheme } from "../../constants";

const { width } = Dimensions.get("screen");


const Comment = ({commentCount}) => {
  const [show, setShow] = useState(false);
  const [comment, setComment] = useState("");
  const [showComment, setShowComment] = useState(false);
  const [count, setCount] = useState(commentCount);


  function sendValues(comment) {
    console.log(comment);
    setCount(count => count + 1);
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

      <View >
        <Input placeholder="Write a comment..."
            style={{marginLeft: -80,
            width: width * 0.835,
            }}
            color = "black"
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
      container: {
        flex: 1,
      },
    button: {
        borderless: true,
        borderColor: '#fff',
        shadowColor: '#fff',
        backgroundColor: '#fff',
        height: 30,
        marginLeft: width*0.25,
        marginTop: -59,
        zIndex : 1,
    },
    submit: {
      position: 'absolute',
      marginLeft: width * 0.585,
      height: 30,
      marginTop: 23
    },
    close: {
      position: 'absolute',
      marginLeft: 315,
      height: 15,
      marginTop: -17
    },
    viewComments: {
      position: 'absolute',
      marginLeft: 190,
      height: 15,
      marginTop: -17
    },
});

export default Comment;
    