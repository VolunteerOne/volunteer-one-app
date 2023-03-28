
import React from "react";
import { Block, Button, Text } from "galio-framework";
import { StyleSheet, Image } from "react-native";

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
        <Block center>
            <Image source={require('../../assets/nucleo icons/svg/thumb-up.svg')} />
            <Button style={styles.button}
            onPress={incrementLikes} >        
            <Text>Likes: {this.state.count}</Text>
          </Button>
          </Block>

        );
        }
        
}
const styles = StyleSheet.create({
    button: {
        borderless: true,
        backgroundColor: '#fff',
    }

});

export default LikeButton;
    