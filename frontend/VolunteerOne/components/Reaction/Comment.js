
import React, { useState }  from "react";
import { Block, Button, Card, Text, Input, theme} from "galio-framework";
import { StyleSheet, TextInput, View, TouchableOpacity } from "react-native";
import Icon from "../Icon";

const Comment = () => {
        const [show, setShow] = useState(false);

        return (
            <View>
            <TouchableOpacity
                style={styles.button}
                onPress={() => setShow(!show)}
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
          <Card style={styles.comment}>
            <TextInput
            placeholder="Write a comment..."
            style={{
                marginTop: -15, paddingLeft: 15,
            width: 360}}
            ></TextInput> 
            </Card>
                : null }
</View>
        );
};
const styles = StyleSheet.create({
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
    },
    comment: {
      backgroundColor: "#FFFFFF",
      borderRadius: 8,
      width: 366,
      marginLeft: -82,
      marginTop: -5,
    }

});

export default Comment;
    