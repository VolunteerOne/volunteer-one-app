
import React from "react";
import { Block, Button, Text } from "galio-framework";
import { StyleSheet, Image } from "react-native";
import { Svg } from "react-native-svg";

class LikeButton extends React.Component {
    state = {
        count: 0
    }

    render() {    
        const incrementLikes = () => {
            console.log("Like button pressed")
            let newCount = this.state.count + 1
            this.setState ({
                count: newCount
            })
            
        }

        return (
            <Button style={styles.button}
            onPress={incrementLikes} >        
            <Text style={styles.titleText}>Likes: {this.state.count}</Text>
          </Button>

        );
        }
        
}
const styles = StyleSheet.create({
    titleText: {
        fontSize: 18,
        color: "#32325D",
      },
    button: {
        borderless: true,
        backgroundColor: '#fff',
        shadowColor: '#fff',
        width: 'auto',
        height: 20,
        marginTop: -30,
        marginLeft: 20,
    }

});

export default LikeButton;
    