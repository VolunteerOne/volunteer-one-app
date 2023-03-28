
import React from "react";
import { Button, Text } from "galio-framework";
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
          <Button
            onPress={incrementLikes}
            >        
            <Image source={require("../../assets/nucleo icons/svg/thumb-up.svg")} />
            <Text>{this.state.count}</Text>
          </Button>
        );
      }
}

export default LikeButton;
