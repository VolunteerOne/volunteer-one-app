
import React from "react";
import { Block, Button, Text } from "galio-framework";
import { StyleSheet, TouchableOpacity, Pressable, View, Dimensions} from "react-native";
import Icon from "../Icon";

const { width } = Dimensions.get("screen");

class LikeButton extends React.Component {
    state = {
        count: this.props.likeCount
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
            <View style={styles.block}>
            <Button style={styles.button}
            onPress={incrementLikes} >      
      <Block flex row style={styles.button2}>
                <Icon
      family="MaterialIcons"
      size={15}
      name="thumb-up"
      color="#32325D"
    />
            {this.state.count == 0 ? 
            <Text style={styles.titleText}>Like</Text>
            :
            <Text style={styles.titleText}>{this.state.count} Likes</Text>  
            }

            </Block>
          </Button>
          </View>
        );
        }
        
}
const styles = StyleSheet.create({
    block: {
        marginLeft: width*0.07,
        zIndex: 1,
        height:33.5
    },
    button2: {
        marginTop: 8
    },
    titleText: {
        fontSize: 15,
        color: "#32325D",
        marginLeft: 5,
        marginTop: -1,
        textAlign: 'right',
    },
    button: {
        borderless: true,
        shadowColor: '#fff',
        flex: 1,
        position: 'absolute',
        bottom: 53,
        width: 90,
        height: 40,
        backgroundColor: '#fff'
    }

});

export default LikeButton;
    