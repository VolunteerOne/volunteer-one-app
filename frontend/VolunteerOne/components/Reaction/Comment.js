
import React, { useState }  from "react";
import { Block, Button, Text, Input, theme} from "galio-framework";
import { StyleSheet, TextInput, View } from "react-native";
import Icon from "../Icon";

const Comment = () => {
        const [show, setShow] = useState(false);

        return (
            <View>
            <Button
                style={styles.button}
                onPress={() => setShow(!show)}
                >
                <Block flex row>
        <Icon
      family="MaterialIcons"
      size={20}
      name="comment"
      color="#32325D"
    />
                <Text style={styles.titleText}>Comment</Text>
                </Block>

          </Button>

          { show ? 
            <Input
            placeholder="Write a comment"
            style={{padding: 10,
                marginLeft: -90,
            width: 360}}
            ></Input> 
                : null }
</View>
        );
};
const styles = StyleSheet.create({
    titleText: {
        fontSize: 18,
        color: "#32325D",
        textAlign: 'right',
        marginLeft: 7,
      },
    button: {
        borderless: true,
        borderColor: '#fff',
        shadowColor: '#fff',
        backgroundColor: '#fff',
        height: 30,
        marginLeft: 150,
        marginTop: -30,

    }

});

export default Comment;
    