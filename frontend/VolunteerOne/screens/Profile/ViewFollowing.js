import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';
import ProfileItem from '../../components/ProfileItem';
import friends from '../../constants/friends';

// ================================= View Friends Page ================================= //

const { width } = Dimensions.get('screen');

class ViewFollowingPage extends React.Component {
  renderFollowing = () => {
    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.notifications}>
      
      {/* map all profiles using loop */}
      <Block flex>
          {friends.map(profile => (
             <ProfileItem item={profile} horizontal following />
          ))}
        </Block>
    
      </ScrollView>
    )
  }

  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderFollowing()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  notifications: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default ViewFollowingPage;
