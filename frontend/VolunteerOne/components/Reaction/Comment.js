
import React, { useState }  from "react";
import { Block, Button, Text, Input, theme} from "galio-framework";
import { StyleSheet, TextInput, View } from "react-native";

const Comment = () => {
        const [show, setShow] = useState(false);

        return (
            <View>
            <Button
                style={styles.button}
                onPress={() => setShow(!show)}>
                <Text style={styles.titleText}>Comment</Text>
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
      },
    button: {
        borderless: true,
        borderColor: '#fff',
        shadowColor: '#fff',
        backgroundColor: '#fff',
        width: 'auto',
        height: 20,
        marginLeft: 160,
        marginTop: -30,

    }

});

export default Comment;
    